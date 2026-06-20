package pnw

import "context"

// AllNations iterates over all pages of nations matching the filter, calling fn
// for each page. Iteration stops when fn returns false or an error occurs.
func (c *Client) AllNations(ctx context.Context, filter NationFilter, fn func([]Nation) bool, fields ...string) error {
	if filter.First == 0 {
		filter.First = 500
	}
	filter.Page = 1
	for {
		page, err := c.Nations(ctx, filter, fields...)
		if err != nil {
			return err
		}
		if !fn(page.Data) || !page.PaginatorInfo.HasMorePages {
			return nil
		}
		filter.Page++
	}
}

// AllAlliances iterates over all pages of alliances matching the filter, calling
// fn for each page. Iteration stops when fn returns false or an error occurs.
func (c *Client) AllAlliances(ctx context.Context, filter AllianceFilter, fn func([]Alliance) bool, fields ...string) error {
	if filter.First == 0 {
		filter.First = 50
	}
	filter.Page = 1
	for {
		page, err := c.Alliances(ctx, filter, fields...)
		if err != nil {
			return err
		}
		if !fn(page.Data) || !page.PaginatorInfo.HasMorePages {
			return nil
		}
		filter.Page++
	}
}

// AllCities iterates over all pages of cities matching the filter, calling fn
// for each page. Iteration stops when fn returns false or an error occurs.
func (c *Client) AllCities(ctx context.Context, filter CityFilter, fn func([]City) bool, fields ...string) error {
	if filter.First == 0 {
		filter.First = 500
	}
	filter.Page = 1
	for {
		page, err := c.Cities(ctx, filter, fields...)
		if err != nil {
			return err
		}
		if !fn(page.Data) || !page.PaginatorInfo.HasMorePages {
			return nil
		}
		filter.Page++
	}
}
