package init

//func NewCmdInit() *cobra.Command {
//	var cmd = &cobra.Command{
//		Use:   "init --node.endpoint=[endpoint]",
//		Short: "init client",
//		Long:  "Get the client ready to use",
//		Args:  cobra.NoArgs,
//		RunE: func(cmd *cobra.Command, args []string) error {
//			ctx := cmd.Context()
//
//			clt, err := kwil_client.New(ctx, config.AppConfig)
//			if err != nil {
//				return err
//			}
//
//			nodeInfo, err := clt.GetNodeInfo(ctx)
//			if err != nil {
//				return err
//			}
//
//			viper.Set("fund.pool_address", nodeInfo.FundingPool)
//			return viper.WriteConfig()
//		},
//	}
//
//	return cmd
//}