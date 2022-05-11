package webscraper

import (
	"real_estate_finder/real_estate_finder/internal/repository/types"
	"testing"
)

func TestBuildURL(t *testing.T) {
	const baseURL = "https://realty.yandex.ru/moskva_i_moskovskaya_oblast/kupit/kvartira/"

	testCases := []struct {
		name          string
		propertyTypes []types.PropertyType
		expected      string
	}{
		{
			name:          "no types selected",
			propertyTypes: nil,
			expected:      baseURL,
		},
		{
			name: "one type",
			propertyTypes: []types.PropertyType{
				types.PropertyType0,
			},
			expected: "https://realty.yandex.ru/moskva_i_moskovskaya_oblast/kupit/kvartira/?roomsTotal=STUDIO",
		},
		{
			name: "two types",
			propertyTypes: []types.PropertyType{
				types.PropertyType0,
				types.PropertyType1,
			},
			expected: "https://realty.yandex.ru/moskva_i_moskovskaya_oblast/kupit/kvartira/?roomsTotal=STUDIO&roomsTotal=1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := buildURL(tc.propertyTypes, 0)
			if actual != tc.expected {
				t.Fatalf("actual != expected")
			}
		})
	}
}
