package apisvc

/*
func (s *Service) PlanSchema(ctx context.Context, req *apipb.PlanSchemaRequest) (*apipb.PlanSchemaResponse, error) {
	planReq := metadata.SchemaRequest{
		Wallet:     req.Wallet,
		Database:   req.Database,
		SchemaData: req.Schema,
	}

	plan, err := s.md.Plan(ctx, planReq)
	if err != nil {
		return nil, err
	}

	changes := make([]*apipb.Change, len(plan.Changes))
	for i, change := range plan.Changes {
		changes[i] = &apipb.Change{
			Cmd:     change.Cmd,
			Comment: change.Comment,
		}
	}

	return &apipb.PlanSchemaResponse{
		Plan: &apipb.Plan{
			Changes: changes,
		},
	}, nil

}

func (s *Service) ApplySchema(ctx context.Context, req *apipb.ApplySchemaRequest) (*apipb.ApplySchemaResponse, error) {
	planReq := metadata.SchemaRequest{
		Wallet:     req.Wallet,
		Database:   req.Database,
		SchemaData: req.Schema,
	}
	err := s.md.Apply(ctx, planReq)
	return &apipb.ApplySchemaResponse{}, err
}

func (s *Service) GetMetadata(ctx context.Context, req *apipb.GetMetadataRequest) (*apipb.GetMetadataResponse, error) {

	mdr := &apipb.Metadata{
		Name: "test",
		Queries: []*apipb.Query{
			{
				Name: "query1",
				Inputs: []*apipb.Param{
					{
						Name: "input1",
						Type: 2, // 2 = string
					},
					{
						Name: "input2",
						Type: 2,
					},
				},
				Outputs: []*apipb.Param{
					{
						Name: "output1",
						Type: 2,
					},
				},
			},
			{
				Name: "wallets:select",
				Inputs: []*apipb.Param{
					{
						Name: "wallet",
						Type: 2,
					},
				},
				Outputs: []*apipb.Param{
					{
						Name: "output1",
						Type: 2,
					},
				},
			},
			{
				Name: "wallets:insert",
				Inputs: []*apipb.Param{
					{
						Name: "wallet",
						Type: 2,
					},
					{
						Name: "wallet_id",
						Type: 2,
					},
				},
				Outputs: []*apipb.Param{
					{
						Name: "output1",
						Type: 2,
					},
				},
			},
			{
				Name: "wallets:update",
				Inputs: []*apipb.Param{
					{
						Name: "wallet",
						Type: 2,
					},
					{
						Name: "balance",
						Type: 2,
					},
				},
				Outputs: []*apipb.Param{
					{
						Name: "output1",
						Type: 2,
					},
				},
			},
			{
				Name: "wallets:delete",
				Inputs: []*apipb.Param{
					{
						Name: "wallet",
						Type: 2,
					},
				},
				Outputs: []*apipb.Param{
					{
						Name: "output1",
						Type: 2,
					},
				},
			},
		},
		Roles: []*apipb.Role{
			{
				Name:    "test",
				Queries: []string{"query1", "query2"},
			},
		},
		Tables: []*apipb.Table{
			{
				Name: "table1",
				Columns: []*apipb.Column{
					{
						Name: "column1",
						Type: 2,
					},
					{
						Name: "column2",
						Type: 2,
					},
				},
			},
		},
	}

	return &apipb.GetMetadataResponse{
		Metadata: mdr,
	}, nil
}

func convertMetadata(meta metadata.Metadata) *apipb.Metadata {
	tables := make([]*apipb.Table, len(meta.Tables))
	for i, table := range meta.Tables {
		tables[i] = convertTable(table)
	}
	queries := make([]*apipb.Query, len(meta.Queries))
	for i, query := range meta.Queries {
		queries[i] = convertQuery(query)
	}

	roles := make([]*apipb.Role, len(meta.Roles))
	for i, role := range meta.Roles {
		roles[i] = convertRole(role)
	}

	return &apipb.Metadata{
		Name:        meta.DbName,
		Tables:      tables,
		Queries:     queries,
		Roles:       roles,
		DefaultRole: meta.DefaultRole,
	}
}

func convertTable(table metadata.Table) *apipb.Table {
	columns := make([]*apipb.Column, len(table.Columns))
	for i, column := range table.Columns {
		columns[i] = convertColumn(column)
	}

	return &apipb.Table{
		Name:    table.Name,
		Columns: columns,
	}
}

func convertColumn(column metadata.Column) *apipb.Column {
	return &apipb.Column{
		Name:  column.Name,
		Type:  convertType(column.Type),
		Arity: convertArity(column.Arity),
	}
}

func convertQuery(query metadata.Query) *apipb.Query {
	inputs := make([]*apipb.Param, len(query.Inputs))
	outputs := make([]*apipb.Param, len(query.Outputs))

	for i, input := range query.Inputs {
		inputs[i] = &apipb.Param{
			Name:  input.Name,
			Type:  convertType(input.Type),
			Arity: convertArity(input.Arity),
		}
	}

	for i, output := range query.Outputs {
		outputs[i] = &apipb.Param{
			Name:  output.Name,
			Type:  convertType(output.Type),
			Arity: convertArity(output.Arity),
		}
	}
	return &apipb.Query{
		Name:    query.Name,
		Inputs:  inputs,
		Outputs: outputs,
	}
}

func convertRole(role metadata.Role) *apipb.Role {
	return &apipb.Role{
		Name:    role.Name,
		Queries: role.Queries,
	}
}

func convertArity(arity metadata.TypeArity) apipb.Arity {
	switch arity {
	case metadata.Optional:
		return apipb.Arity_OPTIONAL
	case metadata.Required:
		return apipb.Arity_REQUIRED
	case metadata.Repeated:
		return apipb.Arity_REPEATED
	default:
		return apipb.Arity_OPTIONAL
	}
}

func convertType(t string) apipb.ParamType {
	switch t {
	case metadata.ScalarString:
		return apipb.ParamType_STRING
	case metadata.ScalarNumber:
		return apipb.ParamType_NUMBER
	case metadata.ScalarBool:
		return apipb.ParamType_BOOL
	case metadata.ScalarDate:
		return apipb.ParamType_DATE
	case metadata.ScalarTime:
		return apipb.ParamType_TIME
	case metadata.ScalarDateTime:
		return apipb.ParamType_DATETIME
	case metadata.ScalarBytes:
		return apipb.ParamType_BYTES
	default:
		return apipb.ParamType_VOID
	}
}
*/
