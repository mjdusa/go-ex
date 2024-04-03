package reflectex_test

import (
	"testing"

	"github.com/mjdusa/go-ext/pkg/reflectex"
	"github.com/stretchr/testify/assert"
)

type Customer struct {
	ID        uint
	ID8       uint8
	ID16      uint16
	ID32      uint32
	ID64      uint64
	Name      string
	TaxExempt bool
}

type Part struct {
	PartNo        string
	Name          string
	UoM           string
	PriceEach32   float32
	PriceEach64   float64
	PriceEachC64  complex64
	PriceEachC128 complex128
}

type Item struct {
	Part       *Part
	Quantity   int
	Quantity8  int8
	Quantity16 int16
	Quantity32 int32
	Quantity64 int64
}

type Order struct {
	ID          uint64
	Description *string
	Customer    Customer
	ItemsSlice  []Item
	ItemsPtr    *[]Item
	ItemsMap    map[string]Item
}

func TestReflector(t *testing.T) {
	description := "This is a description"
	blankOrder := Order{}
	//nolint:lll,nolintlint  // test string
	blankOrderWant := "Order (Struct: 6 fields)\n  Order.ID: 0 (uint64)\n  Order.Customer (Struct: 7 fields)\n    Order.Customer.ID: 0 (uint)\n    Order.Customer.ID8: 0 (uint8)\n    Order.Customer.ID16: 0 (uint16)\n    Order.Customer.ID32: 0 (uint32)\n    Order.Customer.ID64: 0 (uint64)\n    Order.Customer.Name:  (string)\n    Order.Customer.TaxExempt: false (bool)\n  Order.ItemsSlice (Slice: 0 items)\n  Order.ItemsMap (Map: 0 items)\n"
	order := Order{
		ID:          456,
		Customer:    Customer{ID: 56, ID8: 8, ID16: 16, ID32: 32, ID64: 64, Name: "Martha", TaxExempt: true},
		Description: &description,
		ItemsSlice: []Item{
			{Part: &Part{PartNo: "1357", Name: "WIDget", UoM: "Each", PriceEach32: 32.99}, Quantity: 2},
			{Part: &Part{PartNo: "1322", Name: "Gadget", UoM: "Each", PriceEach64: 64.99}, Quantity: 1},
		},
		ItemsPtr: &[]Item{
			{Part: &Part{PartNo: "1357", Name: "WIDget", UoM: "Each", PriceEach64: 64.99}, Quantity: 2},
			{Part: &Part{PartNo: "1322", Name: "Gadget", UoM: "Each", PriceEach32: 32.99}, Quantity: 1},
		},
		ItemsMap: map[string]Item{
			"1357": {
				Part: &Part{
					PartNo:        "1357",
					Name:          "WIDget",
					UoM:           "Each",
					PriceEach32:   32.99,
					PriceEach64:   64.99,
					PriceEachC64:  1 + 2i,
					PriceEachC128: 3 + 4i,
				},
				Quantity: 2,
			},
		},
	}
	//nolint:lll,nolintlint  // test string
	orderWant := "Order (Struct: 6 fields)\n  Order.ID: 456 (uint64)\n  Order.*Description: This is a description (string)\n  Order.Customer (Struct: 7 fields)\n    Order.Customer.ID: 56 (uint)\n    Order.Customer.ID8: 8 (uint8)\n    Order.Customer.ID16: 16 (uint16)\n    Order.Customer.ID32: 32 (uint32)\n    Order.Customer.ID64: 64 (uint64)\n    Order.Customer.Name: Martha (string)\n    Order.Customer.TaxExempt: true (bool)\n  Order.ItemsSlice (Slice: 2 items)\n    Order.ItemsSlice.Item (Struct: 6 fields)\n      Order.ItemsSlice.Item.*Part (Struct: 7 fields)\n        Order.ItemsSlice.Item.*Part.PartNo: 1357 (string)\n        Order.ItemsSlice.Item.*Part.Name: WIDget (string)\n        Order.ItemsSlice.Item.*Part.UoM: Each (string)\n        Order.ItemsSlice.Item.*Part.PriceEach32: 32.9900016784668 (float32)\n        Order.ItemsSlice.Item.*Part.PriceEach64: 0 (float64)\n        Order.ItemsSlice.Item.*Part.PriceEachC64: (0+0i) (complex64)\n        Order.ItemsSlice.Item.*Part.PriceEachC128: (0+0i) (complex128)\n      Order.ItemsSlice.Item.Quantity: 2 (int)\n      Order.ItemsSlice.Item.Quantity8: 0 (int8)\n      Order.ItemsSlice.Item.Quantity16: 0 (int16)\n      Order.ItemsSlice.Item.Quantity32: 0 (int32)\n      Order.ItemsSlice.Item.Quantity64: 0 (int64)\n    Order.ItemsSlice.Item (Struct: 6 fields)\n      Order.ItemsSlice.Item.*Part (Struct: 7 fields)\n        Order.ItemsSlice.Item.*Part.PartNo: 1322 (string)\n        Order.ItemsSlice.Item.*Part.Name: Gadget (string)\n        Order.ItemsSlice.Item.*Part.UoM: Each (string)\n        Order.ItemsSlice.Item.*Part.PriceEach32: 0 (float32)\n        Order.ItemsSlice.Item.*Part.PriceEach64: 64.99 (float64)\n        Order.ItemsSlice.Item.*Part.PriceEachC64: (0+0i) (complex64)\n        Order.ItemsSlice.Item.*Part.PriceEachC128: (0+0i) (complex128)\n      Order.ItemsSlice.Item.Quantity: 1 (int)\n      Order.ItemsSlice.Item.Quantity8: 0 (int8)\n      Order.ItemsSlice.Item.Quantity16: 0 (int16)\n      Order.ItemsSlice.Item.Quantity32: 0 (int32)\n      Order.ItemsSlice.Item.Quantity64: 0 (int64)\n  Order.*ItemsPtr (Slice: 2 items)\n    Order.*ItemsPtr.Item (Struct: 6 fields)\n      Order.*ItemsPtr.Item.*Part (Struct: 7 fields)\n        Order.*ItemsPtr.Item.*Part.PartNo: 1357 (string)\n        Order.*ItemsPtr.Item.*Part.Name: WIDget (string)\n        Order.*ItemsPtr.Item.*Part.UoM: Each (string)\n        Order.*ItemsPtr.Item.*Part.PriceEach32: 0 (float32)\n        Order.*ItemsPtr.Item.*Part.PriceEach64: 64.99 (float64)\n        Order.*ItemsPtr.Item.*Part.PriceEachC64: (0+0i) (complex64)\n        Order.*ItemsPtr.Item.*Part.PriceEachC128: (0+0i) (complex128)\n      Order.*ItemsPtr.Item.Quantity: 2 (int)\n      Order.*ItemsPtr.Item.Quantity8: 0 (int8)\n      Order.*ItemsPtr.Item.Quantity16: 0 (int16)\n      Order.*ItemsPtr.Item.Quantity32: 0 (int32)\n      Order.*ItemsPtr.Item.Quantity64: 0 (int64)\n    Order.*ItemsPtr.Item (Struct: 6 fields)\n      Order.*ItemsPtr.Item.*Part (Struct: 7 fields)\n        Order.*ItemsPtr.Item.*Part.PartNo: 1322 (string)\n        Order.*ItemsPtr.Item.*Part.Name: Gadget (string)\n        Order.*ItemsPtr.Item.*Part.UoM: Each (string)\n        Order.*ItemsPtr.Item.*Part.PriceEach32: 32.9900016784668 (float32)\n        Order.*ItemsPtr.Item.*Part.PriceEach64: 0 (float64)\n        Order.*ItemsPtr.Item.*Part.PriceEachC64: (0+0i) (complex64)\n        Order.*ItemsPtr.Item.*Part.PriceEachC128: (0+0i) (complex128)\n      Order.*ItemsPtr.Item.Quantity: 1 (int)\n      Order.*ItemsPtr.Item.Quantity8: 0 (int8)\n      Order.*ItemsPtr.Item.Quantity16: 0 (int16)\n      Order.*ItemsPtr.Item.Quantity32: 0 (int32)\n      Order.*ItemsPtr.Item.Quantity64: 0 (int64)\n  Order.ItemsMap (Map: 1 items)\n    Order.ItemsMap.Item (Struct: 6 fields)\n      Order.ItemsMap.Item.*Part (Struct: 7 fields)\n        Order.ItemsMap.Item.*Part.PartNo: 1357 (string)\n        Order.ItemsMap.Item.*Part.Name: WIDget (string)\n        Order.ItemsMap.Item.*Part.UoM: Each (string)\n        Order.ItemsMap.Item.*Part.PriceEach32: 32.9900016784668 (float32)\n        Order.ItemsMap.Item.*Part.PriceEach64: 64.99 (float64)\n        Order.ItemsMap.Item.*Part.PriceEachC64: (1+2i) (complex64)\n        Order.ItemsMap.Item.*Part.PriceEachC128: (3+4i) (complex128)\n      Order.ItemsMap.Item.Quantity: 2 (int)\n      Order.ItemsMap.Item.Quantity8: 0 (int8)\n      Order.ItemsMap.Item.Quantity16: 0 (int16)\n      Order.ItemsMap.Item.Quantity32: 0 (int32)\n      Order.ItemsMap.Item.Quantity64: 0 (int64)\n"

	tests := []struct {
		name string
		data interface{}
		want string
	}{
		{
			name: "Test with blank Order",
			data: blankOrder,
			want: blankOrderWant,
		},
		{
			name: "Test with order",
			data: order,
			want: orderWant,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := reflectex.Reflector(tst.data)

			assert.Equal(t, tst.want, got, "name: %s", tst.name)
		})
	}
}
