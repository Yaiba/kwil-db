package sub

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
	"kwil/x"
	"kwil/x/cfgx"
	"kwil/x/svcx/messaging/mx"
	"kwil/x/syncx"
	"sync"
)

func NewChannelBroker(config cfgx.Config) (ChannelBroker, error) {
	cfg, err := mx.NewReceiverConfig(config)
	if err != nil {
		return nil, err
	}

	cb := &channel_broker{
		receiver_assigned: make(chan ReceiverChannel, 32),
		done:              make(chan x.Void),
		pending:           sync.WaitGroup{},
		mu:                sync.Mutex{},
		shutdown:          syncx.NewChan[x.Void](),
		receivers:         make(map[string]map[int32]ReceiverChannel),
		max_poll_records:  100,
	}

	if len(cfg.ConsumerTopics) == 0 {
		return nil, fmt.Errorf("no topics configured")
	}

	if cfg.Group == "" {
		return nil, fmt.Errorf("group is required")
	}

	c, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Brokers...),
		kgo.ClientID(cfg.Client_id),
		kgo.SASL(plain.Auth{User: cfg.User, Pass: cfg.Pwd}.AsMechanism()),
		kgo.Dialer(cfg.Dialer.DialContext),
		kgo.ConsumeTopics(cfg.ConsumerTopics...),
		kgo.AutoCommitMarks(),
		kgo.ConsumerGroup(cfg.Group),
		kgo.OnPartitionsAssigned(cb.handlePartitionsAssigned),
		kgo.OnPartitionsRevoked(cb.handlePartitionsRevoked),
	)

	if err != nil {
		return nil, err
	}

	ctx, fn := context.WithCancel(context.Background())
	cb.ctx = ctx
	cb.cancelFn = fn
	cb.consumer = c

	return cb, nil
}

func NewTransientReceiver(config cfgx.Config) (TransientReceiver, error) {
	cfg, err := mx.NewReceiverConfig(config)
	if err != nil {
		return nil, err
	}

	if len(cfg.ConsumerTopics) != 1 {
		return nil, fmt.Errorf("transient receiver can only be created for a single topic")
	}

	if cfg.Group != "" {
		return nil, fmt.Errorf("transient receiver cannot be used with a consumer group")
	}

	//var adm *kadm.Client
	//{
	//	cl, err := kgo.NewClient(kgo.SeedBrokers(cfg.Brokers...))
	//	if err != nil {
	//		return nil, err
	//	}
	//	adm = kadm.NewClient(cl)
	//}
	//
	//md, err := adm.Metadata(context.Background(), cfg.ConsumerTopics[0])
	//if err != nil {
	//	return nil, err
	//}
	//
	//detail, ok := md.Topics[cfg.ConsumerTopics[0]]
	//if !ok {
	//	return nil, fmt.Errorf("topic (%s) not found", cfg.ConsumerTopics[0])
	//}
	//
	////map[topic]map[paritition_id]Offset
	//partition_count := len(detail.Partitions)

	c, err := kgo.NewClient(
		//kgo.ConsumePartitions(),
		kgo.Dialer(cfg.Dialer.DialContext),
		kgo.SASL(plain.Auth{User: cfg.User, Pass: cfg.Pwd}.AsMechanism()),
		kgo.ConsumeTopics(cfg.ConsumerTopics[0]),
		kgo.SeedBrokers(cfg.Brokers...),
		kgo.ClientID(cfg.Client_id),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtEnd()))

	if err != nil {
		return nil, err
	}

	ctx, fn := context.WithCancel(context.Background())

	return &transient_receiver{
		c,
		cfg.ConsumerTopics[0],
		make(chan MessageIterator, 32), // todo: buffer should be == to partition count
		make(chan x.Void),
		ctx,
		fn,
		cfg.MaxPollRecords,
		&sync.WaitGroup{},
		&sync.Mutex{},
		false,
	}, nil
}