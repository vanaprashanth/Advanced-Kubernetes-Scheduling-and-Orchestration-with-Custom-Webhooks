package mutation

import (
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
)

// injectAnnotations is a container for the mutation injecting annotations into a pod
type injectAnnotations struct {
	Logger logrus.FieldLogger
}

// injectAnnotations implements the podMutator interface
var _ podMutator = (*injectAnnotations)(nil)

// Name returns the struct name
func (se injectAnnotations) Name() string {
	return "inject_annotations"
}

// Mutate returns a new mutated pod according to set annotations
func (se injectAnnotations) Mutate(pod *corev1.Pod) (*corev1.Pod, error) {
	se.Logger = se.Logger.WithField("mutation", se.Name())
	mpod := pod.DeepCopy()

	// build annotations
	annotations := map[string]string{
		"custom-app": "true",
	}

	// inject annotations into pod
	for _, envVar := range annotations {
		se.Logger.Debugf("pod env injected %s", envVar)
		injectAnnotationsIntoPod(mpod, annotations)
	}

	return mpod, nil
}

// injectAnnotationsIntoPod injects annotations into a pod
func injectAnnotationsIntoPod(pod *corev1.Pod, annotations map[string]string) {
	for key, value := range annotations {
		if !HasAnnotations(pod, key, value) {
			pod.Annotations[key] = value
		}
	}
}

// HasAnnotations returns true if annotations exists false otherwise
func HasAnnotations(pod *corev1.Pod, key, value string) bool {
	for annotationKey, annotation := range pod.Annotations {
		if annotationKey == key && annotation == value {
			return true
		}
	}
	return false
}

