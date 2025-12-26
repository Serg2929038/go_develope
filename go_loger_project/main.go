package main

// "context"
// "fmt"
// "log"
// "time"
// "github.com/ClickHouse/clickhouse-go/v2"

func main() {
	// ctx := context.Background()

	// // Подключение к ClickHouse
	// conn, err := clickhouse.Open(&clickhouse.Options{
	// 	Addr: []string{"localhost:9000"},
	// 	Auth: clickhouse.Auth{
	// 		Database: "default",
	// 		Username: "default",
	// 		Password: "123", // твой пароль
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("connect error: %v", err)
	// }

	// if err := conn.Ping(ctx); err != nil {
	// 	log.Fatalf("ping error: %v", err)
	// }

	// // Простой SELECT всех записей из my_table
	// query := `
	// 	SELECT
	// 		user_id,
	// 		message,
	// 		timestamp,
	// 		metric
	// 	FROM my_table
	// 	ORDER BY user_id
	// 	LIMIT 100
	// 	`

	// rows, err := conn.Query(ctx, query)
	// if err != nil {
	// 	log.Fatalf("query error: %v", err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var (
	// 		userID  uint32
	// 		message string
	// 		ts      time.Time
	// 		metric  float32
	// 	)

	// 	if err := rows.Scan(&userID, &message, &ts, &metric); err != nil {
	// 		log.Fatalf("scan error: %v", err)
	// 	}

	// 	fmt.Printf("user_id=%d message=%q timestamp=%s metric=%f\n",
	// 		userID, message, ts.Format(time.RFC3339), metric)
	// }

	// if err := rows.Err(); err != nil {
	// 	log.Fatalf("rows error: %v", err)
	// }
	task18()
	task1()
}
