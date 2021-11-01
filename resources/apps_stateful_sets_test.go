package resources

import (
	v1 "k8s.io/api/core/v1"
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createStatefulSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	setsClient := mocks.NewMockStatefulSetsClient(ctrl)

	setsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&appsv1.StatefulSetList{Items: []appsv1.StatefulSet{fakeStatefulSet(t)}}, nil,
	)
	return client.Services{
		StatefulSets: setsClient,
	}
}

func fakeStatefulSet(t *testing.T) appsv1.StatefulSet {
	var rs appsv1.StatefulSet
	if err := faker.FakeDataSkipFields(&rs, []string{"Spec"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&rs.Spec, []string{"PodManagementPolicy", "Selector", "Template", "VolumeClaimTemplates"}); err != nil {
		t.Fatal(err)
	}
	rs.Spec.PodManagementPolicy = "test"
	rs.Spec.VolumeClaimTemplates = []v1.PersistentVolumeClaim{*fakePersistentVolumeClaim(t)}
	rs.Spec.Selector = fakeSelector(t)
	rs.Spec.Template = fakePodTemplateSpec(t)
	return rs
}

func TestAppsStatefulSets(t *testing.T) {
	k8sTestHelper(t, AppsStatefulSets(), createStatefulSets)
}
