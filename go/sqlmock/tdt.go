// ref: https://github.com/jvalecillos/sqlmock-examples
package main

func Test_example(t *testing.T) {

	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
	}{
		{
			name: "my test case",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO test \(key, value\) VALUES \(1,1\)`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "my test case",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("INSERT INTO").
					WillBeClosed().
					ExpectExec().
					WithArgs("test title", "test body").
					WillReturnError(fmt.Errorf("DB error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// db and mock for current iteration
			db, mock, _ := sqlmock.New()
			defer db.Close()
			// set mock expectations
			tt.mockClosure(mock)
						
			// ...

			// Checking all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}