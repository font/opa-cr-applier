/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ModelAccuracySpec defines the desired state of ModelAccuracy
type ModelAccuracySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Accuracy of the model as a decimal e.g. 0.9577 for 95.77%.
	Accuracy string `json:"accuracy,omitempty"`
}

// ModelAccuracyStatus defines the observed state of ModelAccuracy
type ModelAccuracyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// ModelAccuracy is the Schema for the modelaccuracies API
type ModelAccuracy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ModelAccuracySpec   `json:"spec,omitempty"`
	Status ModelAccuracyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelAccuracyList contains a list of ModelAccuracy
type ModelAccuracyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ModelAccuracy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ModelAccuracy{}, &ModelAccuracyList{})
}
