// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSenderAuthenticationResource(t *testing.T) {
	resourceName := "sendgrid_sender_authentication.test"

	domain := fmt.Sprintf("test-acc-%s.com", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccSenderAuthenticationResourceConfig(domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "domain", domain),
					resource.TestCheckResourceAttr(resourceName, "valid", "false"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "legacy", "false"),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccSenderAuthenticationResourceConfig(domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "domain", domain),
					resource.TestCheckResourceAttr(resourceName, "valid", "false"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "legacy", "false"),
				),
			},
		},
	})
}

func testAccSenderAuthenticationResourceConfig(domain string) string {
	return fmt.Sprintf(`
resource "sendgrid_sender_authentication" "test" {
  domain = "%[1]s"
}
`, domain)
}

// TestIsValidDkimSelector tests the isValidDkimSelector function
func TestIsValidDkimSelector(t *testing.T) {
	tests := []struct {
		name     string
		selector string
		want     bool
	}{
		{
			name:     "Valid selector - lowercase",
			selector: "abc",
			want:     true,
		},
		{
			name:     "Valid selector - uppercase",
			selector: "XYZ",
			want:     true,
		},
		{
			name:     "Valid selector - alphanumeric",
			selector: "A1b",
			want:     true,
		},
		{
			name:     "Invalid selector - too short",
			selector: "ab",
			want:     false,
		},
		{
			name:     "Invalid selector - too long",
			selector: "abcd",
			want:     false,
		},
		{
			name:     "Invalid selector - special characters",
			selector: "a$b",
			want:     false,
		},
		{
			name:     "Invalid selector - empty",
			selector: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidDkimSelector(tt.selector)
			if got != tt.want {
				t.Errorf("isValidDkimSelector(%s) = %v, want %v", tt.selector, got, tt.want)
			}
		})
	}
}
