package datasource

import (
	"fmt"
	"testing"
)

func TestExecuteSimple(t *testing.T) {
	//ctx := context.Background()

	t.Run("Simple Test", func(t *testing.T) {
		// runner := &MockRunner{
		// 	testDataPath: "simple.csv",
		// }

		// dr := ExecuteQuery(ctx, models.QueryModel{}, runner)

		// if dr.Error != nil {
		// 	t.Fatal(dr.Error)
		// }

		// if len(dr.Frames) != 1 {
		// 	t.Fatalf("Expected 1 frame")
		// }

		// if dr.Frames[0].Name != "test" {
		// 	t.Fatalf("Frame must match _measurement column. Expected [%s] Got [%s]", "test", dr.Frames[0].Name)
		// }

		// if len(dr.Frames[0].Fields[1].Labels) != 2 {
		// 	t.Fatalf("Error parsing labels. Expected [%d] Got [%d]", 2, len(dr.Frames[0].Fields[1].Labels))
		// }

		// if dr.Frames[0].Fields[0].Name != "Time" {
		// 	t.Fatalf("Error parsing fields. Field 1 should always be time. Got name [%s]", dr.Frames[0].Fields[0].Name)
		// }

		// st, _ := dr.Frames[0].StringTable(-1, -1)
		// fmt.Println(st)
		fmt.Println("----------------------")
	})
}
