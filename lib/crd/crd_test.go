/*
Copyright 2022. projectsveltos.io. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package crd_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/projectsveltos/libsveltos/lib/crd"
)

var _ = Describe("CRD", func() {
	It("Gets the Classifier CustomResourceDefinition", func() {
		yaml := crd.GetClassifierCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_classifiers.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the ClassifierReport CustomResourceDefinition", func() {
		yaml := crd.GetClassifierReportCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_classifierreports.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the DebuggingConfiguration CustomResourceDefinition", func() {
		yaml := crd.GetDebuggingConfigurationCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_debuggingconfigurations.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the AccessRequest CustomResourceDefinition", func() {
		yaml := crd.GetAccessRequestCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_accessrequests.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the SveltosCluster CustomResourceDefinition", func() {
		yaml := crd.GetSveltosClusterCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_sveltosclusters.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the ResourceSummary CustomResourceDefinition", func() {
		yaml := crd.GetResourceSummaryCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_resourcesummaries.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the RoleRequest CustomResourceDefinition", func() {
		yaml := crd.GetRoleRequestCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_rolerequests.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the ClusterHealthCheck CustomResourceDefinition", func() {
		yaml := crd.GetClusterHealthCheckCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_clusterhealthchecks.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the HealthCheck CustomResourceDefinition", func() {
		yaml := crd.GetHealthCheckCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_healthchecks.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the HealthCheckReport CustomResourceDefinition", func() {
		yaml := crd.GetHealthCheckReportCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_healthcheckreports.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the EventSource CustomResourceDefinition", func() {
		yaml := crd.GetEventSourceCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_eventsources.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the EventReport CustomResourceDefinition", func() {
		yaml := crd.GetEventReportCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_eventreports.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the Reloader CustomResourceDefinition", func() {
		yaml := crd.GetReloaderCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_reloaders.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the ReloaderReport CustomResourceDefinition", func() {
		yaml := crd.GetReloaderReportCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_reloaderreports.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the ClusterSet CustomResourceDefinition", func() {
		yaml := crd.GetClusterSetCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_clustersets.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})

	It("Gets the Set CustomResourceDefinition", func() {
		yaml := crd.GetSetCRDYAML()

		filename := "../../config/crd/bases/lib.projectsveltos.io_sets.yaml"
		currentFile, err := os.ReadFile(filename)
		Expect(err).To(BeNil())

		Expect(string(yaml)).To(Equal(string(currentFile)))
	})
})
