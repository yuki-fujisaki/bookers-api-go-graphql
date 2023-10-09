package bookers

import (
	"bookers/ent"
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func Example_Book() {
	// インメモリーのSQLiteデータベースを持つent.Clientを作成します。
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// 自動マイグレーションツールを実行して、すべてのスキーマリソースを作成します。
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	book1, err := client.Book.Create().Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a book: %v", err)
	}
	fmt.Println(book1)
}
