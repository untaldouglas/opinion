package opinion

import (
	"context"
	"fmt"
	"log"
	"opinion/ent"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func Example_Opinion() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Output:
	// ...

	task1, err := client.Opinion.
		Create().
		SetAsunto("Golang es mejor que Python").
		SetContenido("Por economía, rendimiento, y comunidad").
		Save(ctx)
	if err != nil {
		log.Fatalf("Error al crear una opinion: %v", err)
	}
	fmt.Printf("%d: %q %q\n", task1.ID, task1.Asunto, task1.Contenido)

	// Output:
	// 1: "Golang es mejor que Python" "Por economía, rendimiento, y comunidad"

}
