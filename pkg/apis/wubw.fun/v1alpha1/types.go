package v1alpha1

import metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type DemoSpec struct {
	Param1 string `json:"param1,omitempty"`
}

type DemoStatus struct {
	Ok bool `json:"ok,omitempty"`
}

type Demo struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoSpec   `json:"spec"`
	Status DemoStatus `json:"status"`
}

type DemoList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`
	Items           []Demo `json:"items"`
}
