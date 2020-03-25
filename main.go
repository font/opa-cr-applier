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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	aiv1alpha1 "github.com/okd-apps/opa-cr-applier/api/v1alpha1"
	yaml "gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	setupLog = ctrl.Log.WithName("main")
)

//func init() {
//	_ = clientgoscheme.AddToScheme(scheme)
//
//	_ = aiv1alpha1.AddToScheme(scheme)
//	// +kubebuilder:scaffold:scheme
//}

func main() {
	var metricsFile, outputFile string
	flag.StringVar(&metricsFile, "metrics-file", "", "The file name containing the accuracy metrics of an AI model.")
	flag.StringVar(&outputFile, "output-file", "", "The file name used to output the AI model custom resource.")
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	if metricsFile == "" {
		setupLog.Info("Required argument", "metrics-file", metricsFile)
		os.Exit(1)
	} else if outputFile == "" {
		setupLog.Info("Required argument", "output-file", outputFile)
		os.Exit(1)
	}

	jsonData, err := ioutil.ReadFile(metricsFile)
	if err != nil {
		setupLog.Error(err, "Error reading file", "metrics-file", metricsFile)
		os.Exit(1)
	}

	var jsonMap map[string]interface{}

	if err := json.Unmarshal(jsonData, &jsonMap); err != nil {
		setupLog.Error(err, "Error in JSON Unmarshal", "json", jsonData)
		os.Exit(1)
	}

	accuracy := jsonMap["accuracy"].(float64)
	fmt.Printf("accuracy = %v\n", accuracy)

	modelAccuracy := &aiv1alpha1.ModelAccuracy{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ModelAccuracy",
			APIVersion: "ai.ifontlabs.com/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "modelaccuracy-",
		},
		Spec: aiv1alpha1.ModelAccuracySpec{
			Accuracy: int64(accuracy * 100),
		},
	}
	jsonData, err = json.Marshal(modelAccuracy)
	if err != nil {
		setupLog.Error(err, "error marshaling into JSON")
		os.Exit(1)
	}

	var jsonObj interface{}
	err = yaml.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		setupLog.Error(err, "error in yaml Unmarshal")
		os.Exit(1)
	}

	// Marshal this object into YAML.
	yamlObj, err := yaml.Marshal(jsonObj)
	if err != nil {
		setupLog.Error(err, "Error marshalling into YAML")
		os.Exit(1)
	}
	fmt.Printf("%v\n", string(yamlObj))
	err = ioutil.WriteFile(outputFile, yamlObj, 0644)
	if err != nil {
		setupLog.Error(err, "Error writing YAML object to file", "output-file", outputFile, "yaml", yamlObj)
		os.Exit(1)
	}
}
