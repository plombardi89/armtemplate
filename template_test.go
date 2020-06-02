package armtemplate_test

import (
	"fmt"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/plombardi89/armtemplate"
	"github.com/plombardi89/armtemplate/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTemplate_Render(t *testing.T) {
	testCases := []struct {
		template func() armtemplate.Template
		expected string
	}{
		{
			expected: testutil.Data(t, "template_new.json"),
			template: armtemplate.New,
		},
		{
			expected: testutil.Data(t, "template_vars_0.json"),
			template: templateVars0,
		},
		{
			expected: testutil.Data(t, "real/linux_vm.json"),
			template: templateLinuxVM,
		},
	}

	ja := jsonassert.New(t)

	for _, tc := range testCases {
		rendered, err := tc.template().Render()
		if assert.NoError(t, err) {
			fmt.Println(rendered)
			ja.Assertf(rendered, tc.expected)
		}
	}
}

func templateLinuxVM() armtemplate.Template {
	tmpl := armtemplate.New()

	tmpl.WithContentVersion("1.0.0.0")

	tmpl.WithParameters(func(p armtemplate.Parameters) {
		p.String("projectName").
			Description("Specifies a name for generating resource names.")

		p.String("location").
			Default("[resourceGroup().location]").
			Description("Specifies the location for all resources.")

		p.String("adminUsername").
			Description("Specifies a username for the Virtual Machine.")

		p.String("adminPublicKey").
			Description(`Specifies the SSH rsa public key file as a string. Use "ssh-keygen -t rsa -b 2048" to generate your SSH key pairs.`)
	})

	tmpl.WithVariables(func(v armtemplate.Variables) {
		v.String("vNetName", "[concat(parameters('projectName'), '-vnet')]")
		v.String("vNetAddressPrefixes", "10.0.0.0/16")
		v.String("vNetSubnetName", "default")
		v.String("vNetSubnetAddressPrefix", "10.0.0.0/24")
		v.String("vmName", "[concat(parameters('projectName'), '-vm')]")
		v.String("publicIPAddressName", "[concat(parameters('projectName'), '-ip')]")
		v.String("networkInterfaceName", "[concat(parameters('projectName'), '-nic')]")
		v.String("networkSecurityGroupName", "[concat(parameters('projectName'), '-nsg')]")
		v.String("networkSecurityGroupName2", "[concat(variables('vNetSubnetName'), '-nsg')]")
	})

	tmpl.WithResources(func(r armtemplate.Resources) {
		r.Add(
			armtemplate.Resource{
				Type:       "Microsoft.Network/networkSecurityGroups",
				APIVersion: "2018-11-01",
				Name:       "[variables('networkSecurityGroupName')]",
				Location:   "[parameters('location')]",
				Properties: armtemplate.Properties{
					"securityRules": []armtemplate.Properties{
						{
							"name": "ssh_rule",
							"properties": armtemplate.Properties{
								"description":              "Locks inbound down to ssh default port 22.",
								"protocol":                 "Tcp",
								"sourcePortRange":          "*",
								"sourceAddressPrefix":      "*",
								"destinationPortRange":     "22",
								"destinationAddressPrefix": "*",
								"access":                   "Allow",
								"priority":                 123,
								"direction":                "Inbound",
							},
						},
					},
				},
			})

		r.Add(armtemplate.Resource{
			Type:       "Microsoft.Network/publicIPAddresses",
			APIVersion: "2018-11-01",
			Name:       "[variables('publicIPAddressName')]",
			Location:   "[parameters('location')]",
			Properties: armtemplate.Properties{
				"publicIPAllocationMethod": "Dynamic",
			},
			SKU: &armtemplate.ResourceSKU{Name: "Basic"},
		})

		r.Add(armtemplate.Resource{
			Comment:    "Simple Network Security Group for subnet [variables('vNetSubnetName')]",
			Type:       "Microsoft.Network/networkSecurityGroups",
			APIVersion: "2019-08-01",
			Name:       "[variables('networkSecurityGroupName2')]",
			Location:   "[parameters('location')]",
			Properties: armtemplate.Properties{
				"securityRules": []armtemplate.Properties{
					{
						"name": "default-allow-22",
						"properties": armtemplate.Properties{
							"protocol":                 "Tcp",
							"sourcePortRange":          "*",
							"sourceAddressPrefix":      "*",
							"destinationPortRange":     "22",
							"destinationAddressPrefix": "*",
							"access":                   "Allow",
							"priority":                 1000,
							"direction":                "Inbound",
						},
					},
				},
			},
		},
		)

		r.Add(armtemplate.Resource{
			Type:       "Microsoft.Network/virtualNetworks",
			APIVersion: "2018-11-01",
			Name:       "[variables('vNetName')]",
			Location:   "[parameters('location')]",
			DependsOn: armtemplate.DependsOn{
				"[resourceId('Microsoft.Network/networkSecurityGroups', variables('networkSecurityGroupName2'))]",
			},
			Properties: armtemplate.Properties{
				"addressSpace": armtemplate.Properties{
					"addressPrefixes": []string{"[variables('vNetAddressPrefixes')]"},
				},
				"subnets": []armtemplate.Properties{
					{
						"name": "[variables('vNetSubnetName')]",
						"properties": armtemplate.Properties{
							"addressPrefix": "[variables('vNetSubnetAddressPrefix')]",
							"networkSecurityGroup": armtemplate.Properties{
								"id": "[resourceId('Microsoft.Network/networkSecurityGroups', variables('networkSecurityGroupName2'))]",
							},
						},
					},
				},
			},
		})

		r.Add(armtemplate.Resource{
			Type:       "Microsoft.Network/networkInterfaces",
			APIVersion: "2018-11-01",
			Name:       "[variables('networkInterfaceName')]",
			Location:   "[parameters('location')]",
			DependsOn: armtemplate.DependsOn{
				"[resourceId('Microsoft.Network/publicIPAddresses', variables('publicIPAddressName'))]",
				"[resourceId('Microsoft.Network/virtualNetworks', variables('vNetName'))]",
				"[resourceId('Microsoft.Network/networkSecurityGroups', variables('networkSecurityGroupName'))]",
			},
			Properties: armtemplate.Properties{
				"ipConfigurations": []armtemplate.Properties{
					{
						"name": "ipconfig1",
						"properties": armtemplate.Properties{
							"privateIPAllocationMethod": "Dynamic",
							"publicIPAddress": armtemplate.Properties{
								"id": "[resourceId('Microsoft.Network/publicIPAddresses', variables('publicIPAddressName'))]",
							},
							"subnet": armtemplate.Properties{
								"id": "[resourceId('Microsoft.Network/virtualNetworks/subnets', variables('vNetName'), variables('vNetSubnetName'))]",
							},
						},
					},
				},
			},
		})

		r.Add(armtemplate.Resource{
			Type:       "Microsoft.Compute/virtualMachines",
			APIVersion: "2018-10-01",
			Name:       "[variables('vmName')]",
			Location:   "[parameters('location')]",
			DependsOn: armtemplate.DependsOn{
				"[resourceId('Microsoft.Network/networkInterfaces', variables('networkInterfaceName'))]",
			},
			Properties: armtemplate.Properties{
				"hardwareProfile": armtemplate.Properties{
					"vmSize": "Standard_D2s_v3",
				},
				"osProfile": armtemplate.Properties{
					"computerName":  "[variables('vmName')]",
					"adminUsername": "[parameters('adminUsername')]",
					"linuxConfiguration": armtemplate.Properties{
						"disablePasswordAuthentication": true,
						"ssh": armtemplate.Properties{
							"publicKeys": []armtemplate.Properties{
								{
									"path":    "[concat('/home/', parameters('adminUsername'), '/.ssh/authorized_keys')]",
									"keyData": "[parameters('adminPublicKey')]",
								},
							},
						},
					},
				},
				"storageProfile": armtemplate.Properties{
					"imageReference": armtemplate.Properties{
						"publisher": "Canonical",
						"offer":     "UbuntuServer",
						"sku":       "18.04-LTS",
						"version":   "latest",
					},
					"osDisk": armtemplate.Properties{
						"createOption": "fromImage",
					},
				},
				"networkProfile": armtemplate.Properties{
					"networkInterfaces": []armtemplate.Properties{
						{
							"id": "[resourceId('Microsoft.Network/networkInterfaces', variables('networkInterfaceName'))]",
						},
					},
				},
			},
		})
	})

	tmpl.WithOutputs(func(o armtemplate.Outputs) {
		o.String("adminUsername", o.WithValue("[parameters('adminUsername')]"))
	})

	return tmpl
}

func templateVars0() armtemplate.Template {
	tmpl := armtemplate.New()
	vars := tmpl.Variables()

	vars.String("var1", "myVariable")
	vars.IntArray("var2", []int{1, 2, 3, 4})
	vars.Reference("var3", "var1")
	vars.Map("var4", map[string]interface{}{
		"property1": "value1",
		"property2": "value2",
	})
	vars.Copy("var5", []armtemplate.Copy{
		{
			Name:  "disks",
			Count: 5,
			Input: map[string]interface{}{
				"name":       "[concat('myDataDisk', copyIndex('disks', 1))]",
				"diskSizeGB": "1",
				"diskIndex":  "[copyIndex('disks')]",
			},
		},
		{
			Name:  "diskNames",
			Count: 5,
			Input: "[concat('myDataDisk', copyIndex('diskNames', 1))]",
		},
	})

	vars.Copy("copy", []armtemplate.Copy{
		{
			Name:  "var6",
			Count: 5,
			Input: map[string]interface{}{
				"name":       "[concat('oneDataDisk', copyIndex('var6', 1))]",
				"diskSizeGB": "1",
				"diskIndex":  "[copyIndex('var6')]",
			},
		},
		{
			Name:  "var7",
			Count: 3,
			Input: map[string]interface{}{
				"name":       "[concat('twoDataDisk', copyIndex('var7', 1))]",
				"diskSizeGB": "1",
				"diskIndex":  "[copyIndex('var7')]",
			},
		},
		{
			Name:  "var8",
			Count: 4,
			Input: "[concat('stringValue', copyIndex('var8'))]",
		},
		{
			Name:  "var9",
			Count: 4,
			Input: "[copyIndex('var9')]",
		},
	})

	return tmpl
}
