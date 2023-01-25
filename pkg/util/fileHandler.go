package util

import (
	"fmt"
	"os"

	skiperator "github.com/kartverket/skiperator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/yaml"
)

func ReadApplicationFromFile(path string) (skiperator.Application, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	scheme := runtime.NewScheme()
	_ = skiperator.AddToScheme(scheme)
	_ = clientgoscheme.AddToScheme(scheme)

	if err != nil {
		fmt.Printf("%#v", err)
	}

	decode := serializer.NewCodecFactory(scheme).UniversalDeserializer().Decode

	obj, _, err := decode(b, nil, nil)
	if err != nil {
		// TODO Error handling if cannot decode

		return skiperator.Application{}, err
	}

	application := obj.(*skiperator.Application)

	return *application, nil
}

func WriteApplicationToFile(filename string, application skiperator.Application) error {
	applicationYaml, _ := yaml.Marshal(application)

	filenameFull := fmt.Sprintf("%s.yaml", filename)

	err := os.WriteFile(filenameFull, []byte(applicationYaml), 0644)
	if err != nil {
		// TODO Error handling if cannot decode
		return err
	}

	return nil
}

func printScheme(scheme runtime.Scheme) {
	tmap := scheme.AllKnownTypes()
	for key, value := range tmap {
		fmt.Printf("GroupVersionKind : %v, Type : %v\n", key, value)
	}
}
