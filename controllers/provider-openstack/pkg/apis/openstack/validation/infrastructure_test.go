// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation_test

import (
	apisopenstack "github.com/gardener/gardener-extensions/controllers/provider-openstack/pkg/apis/openstack"
	. "github.com/gardener/gardener-extensions/controllers/provider-openstack/pkg/apis/openstack/validation"

	. "github.com/gardener/gardener/pkg/utils/validation/gomega"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

var _ = Describe("InfrastructureConfig validation", func() {
	var (
		floatingPoolName1 = "foo"

		constraints = apisopenstack.Constraints{
			FloatingPools: []apisopenstack.FloatingPool{
				{Name: floatingPoolName1},
			},
		}

		infrastructureConfig *apisopenstack.InfrastructureConfig

		nodes       = "10.250.0.0/16"
		invalidCIDR = "invalid-cidr"
	)

	BeforeEach(func() {
		infrastructureConfig = &apisopenstack.InfrastructureConfig{
			FloatingPoolName: floatingPoolName1,
			Networks: apisopenstack.Networks{
				Router: &apisopenstack.Router{
					ID: "hugo",
				},
				Worker: "10.250.0.0/16",
			},
		}
	})

	Describe("#ValidateInfrastructureConfig", func() {
		It("should forbid invalid floating pool name configuration", func() {
			infrastructureConfig.FloatingPoolName = ""

			errorList := ValidateInfrastructureConfig(infrastructureConfig, constraints, &nodes)

			Expect(errorList).To(ConsistOfFields(Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("floatingPoolName"),
			}))
		})

		Context("CIDR", func() {
			It("should forbid invalid workers CIDR", func() {
				infrastructureConfig.Networks.Worker = invalidCIDR

				errorList := ValidateInfrastructureConfig(infrastructureConfig, constraints, &nodes)

				Expect(errorList).To(ConsistOfFields(Fields{
					"Type":   Equal(field.ErrorTypeInvalid),
					"Field":  Equal("networks.worker"),
					"Detail": Equal("invalid CIDR address: invalid-cidr"),
				}))
			})

			It("should forbid workers CIDR which are not in Nodes CIDR", func() {
				infrastructureConfig.Networks.Worker = "1.1.1.1/32"

				errorList := ValidateInfrastructureConfig(infrastructureConfig, constraints, &nodes)

				Expect(errorList).To(ConsistOfFields(Fields{
					"Type":   Equal(field.ErrorTypeInvalid),
					"Field":  Equal("networks.worker"),
					"Detail": Equal(`must be a subset of "" ("10.250.0.0/16")`),
				}))
			})

			It("should forbid non canonical CIDRs", func() {
				nodeCIDR := "10.250.0.3/16"

				infrastructureConfig.Networks.Worker = "10.250.3.8/24"

				errorList := ValidateInfrastructureConfig(infrastructureConfig, constraints, &nodeCIDR)
				Expect(errorList).To(HaveLen(1))

				Expect(errorList).To(ConsistOfFields(Fields{
					"Type":   Equal(field.ErrorTypeInvalid),
					"Field":  Equal("networks.worker"),
					"Detail": Equal("must be valid canonical CIDR"),
				}))
			})
		})
	})

	Describe("#ValidateInfrastructureConfigUpdate", func() {
		It("should return no errors for an unchanged config", func() {
			Expect(ValidateInfrastructureConfigUpdate(infrastructureConfig, infrastructureConfig, constraints, &nodes)).To(BeEmpty())
		})

		It("should forbid changing the network section", func() {
			newInfrastructureConfig := infrastructureConfig.DeepCopy()
			newInfrastructureConfig.Networks.Router = &apisopenstack.Router{ID: "name"}

			errorList := ValidateInfrastructureConfigUpdate(infrastructureConfig, newInfrastructureConfig, constraints, &nodes)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("networks"),
			}))))
		})
	})
})
