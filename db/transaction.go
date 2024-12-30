package db

// ExecTx executes a function within a database transaction
// func (store *Store) execTx(ctx context.Context, fn func(*repository.Repository) error) error {
// 	tx, err := store.connPool.Begin(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	r := repository.InitRepo(tx)
// 	err = fn(r)
// 	if err != nil {
// 		if rbErr := tx.Rollback(ctx); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit(ctx)
// }
