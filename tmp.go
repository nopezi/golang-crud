// Update updates Transaction
// UpdateToExecuted
func (r TransactionRepository) Inquiry(Transaction models.Transaction) (bool, error) {
	transaction := requests.TransactionRequest{
		Appname:       Transaction.Appname,
		Data:          Transaction.Data,
		Prefix:        Transaction.Prefix,
		ExpiredDate:   Transaction.ExpiredDate,
		ReferenceCode: Transaction.ReferenceCode,
		Status:        Transaction.Status,
	}

	bdy, err := json.Marshal(transaction)
	if err != nil {
		return false, fmt.Errorf("insert: marshall: %w", err)
	}

	req := esapi.UpdateRequest{
		Index:      Transaction.IndexName(),
		DocumentID: Transaction.Id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	res, err := req.Do(ctx, r.elastic.Client)
	if err != nil {
		return false, fmt.Errorf("insert: request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return false, fmt.Errorf("insert: response: %s", res.String())
	}

	return true, err
}