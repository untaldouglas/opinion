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

	opinion1, err := client.Opinion.
		Create().
		SetAsunto("Golang es mejor que Python").
		SetContenido("Por economía, rendimiento, y comunidad").
		Save(ctx)
	if err != nil {
		log.Fatalf("Error al crear una opinion: %v", err)
	}
	fmt.Printf("%d: %q %q\n", opinion1.ID, opinion1.Asunto, opinion1.Contenido)
	// Output:
	// 1: "Golang es mejor que Python" "Por economía, rendimiento, y comunidad"

	opinion2, err := client.Opinion.
		Create().
		SetAsunto("Bayer Munich es mejor que PSG").
		SetContenido("Será que se confirma en la champions !!??").
		Save(ctx)
	if err != nil {
		log.Fatalf("Error al crear una opinion: %v", err)
	}
	if err := opinion2.Update().SetParent(opinion1).Exec(ctx); err != nil {
		log.Fatalf("error vinculando opinion2 a su registro padre: %v", err)
	}
	fmt.Printf("%d: %q %q\n", opinion2.ID, opinion2.Asunto, opinion2.Contenido)

	// Query todas las opiniones.
	items, err := client.Opinion.Query().All(ctx)
	if err != nil {
		log.Fatalf("error al consultar todos las opiniones: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Asunto)
	}
	// Output:
	// 1: "Golang es mejor que Python" "Por economía, rendimiento, y comunidad"
	// 2: "Bayer Munich es mejor que PSG" "Será que se confirma en la champions !!??"
	// 1: "Golang es mejor que Python"
	// 2: "Bayer Munich es mejor que PSG"
}
