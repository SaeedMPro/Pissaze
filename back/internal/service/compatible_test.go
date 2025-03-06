package service

import (
	"testing"

	"github.com/pissaze/internal/models"
)

func TestIntersectionSlices(t *testing.T) {
    product := func(id int) models.Product {
        return models.Product{ID: id}
    }

    tests := []struct {
        name     string
        base     []models.Product
        new      []models.Product
        expected []models.Product
    }{
        {
            name:     "Both slices are empty",
            base:     []models.Product{},
            new:      []models.Product{},
            expected: []models.Product{},
        },
        {
            name:     "Base slice is empty",
            base:     []models.Product{},
            new:      []models.Product{product(1), product(2)},
            expected: []models.Product{product(1), product(2)},
        },
        {
            name:     "New slice is empty",
            base:     []models.Product{product(1), product(2)},
            new:      []models.Product{},
            expected: []models.Product{},
        },
        {
            name:     "No intersection",
            base:     []models.Product{product(1), product(2)},
            new:      []models.Product{product(3), product(4)},
            expected: []models.Product{},
        },
        {
            name:     "Partial intersection",
            base:     []models.Product{product(1), product(2), product(3)},
            new:      []models.Product{product(2), product(3), product(4)},
            expected: []models.Product{product(2), product(3)},
        },
        {
            name:     "Full intersection",
            base:     []models.Product{product(1), product(2), product(3)},
            new:      []models.Product{product(1), product(2), product(3)},
            expected: []models.Product{product(1), product(2), product(3)},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := intersectionSlices(tt.base, tt.new)

            if len(result) != len(tt.expected) {
                t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
                return
            }

            if !areSlicesEqual(result, tt.expected){
				t.Error("Expected ", tt.expected, "but got ", result , "\n")
                return
			}
        })
    }
}

//------------------ helpers --------------------
func areSlicesEqual(a, b []models.Product) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i].ID != b[i].ID {
            return false
        }
    }
    return true
}