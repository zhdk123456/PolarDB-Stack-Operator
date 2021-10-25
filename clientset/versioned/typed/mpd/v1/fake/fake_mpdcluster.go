/*
*Copyright (c) 2019-2021, Alibaba Group Holding Limited;
*Licensed under the Apache License, Version 2.0 (the "License");
*you may not use this file except in compliance with the License.
*You may obtain a copy of the License at

*   http://www.apache.org/licenses/LICENSE-2.0

*Unless required by applicable law or agreed to in writing, software
*distributed under the License is distributed on an "AS IS" BASIS,
*WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*See the License for the specific language governing permissions and
*limitations under the License.
 */

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	mpdv1 "github.com/ApsaraDB/PolarDB-Stack-Operator/apis/mpd/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMPDClusters implements MPDClusterInterface
type FakeMPDClusters struct {
	Fake *FakeMpdV1
	ns   string
}

var mpdclustersResource = schema.GroupVersionResource{Group: "mpd", Version: "v1", Resource: "mpdclusters"}

var mpdclustersKind = schema.GroupVersionKind{Group: "mpd", Version: "v1", Kind: "MPDCluster"}

// Get takes name of the mPDCluster, and returns the corresponding mPDCluster object, and an error if there is any.
func (c *FakeMPDClusters) Get(name string, options v1.GetOptions) (result *mpdv1.MPDCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mpdclustersResource, c.ns, name), &mpdv1.MPDCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mpdv1.MPDCluster), err
}

// List takes label and field selectors, and returns the list of MPDClusters that match those selectors.
func (c *FakeMPDClusters) List(opts v1.ListOptions) (result *mpdv1.MPDClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mpdclustersResource, mpdclustersKind, c.ns, opts), &mpdv1.MPDClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mpdv1.MPDClusterList{ListMeta: obj.(*mpdv1.MPDClusterList).ListMeta}
	for _, item := range obj.(*mpdv1.MPDClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mPDClusters.
func (c *FakeMPDClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mpdclustersResource, c.ns, opts))

}

// Create takes the representation of a mPDCluster and creates it.  Returns the server's representation of the mPDCluster, and an error, if there is any.
func (c *FakeMPDClusters) Create(mPDCluster *mpdv1.MPDCluster) (result *mpdv1.MPDCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mpdclustersResource, c.ns, mPDCluster), &mpdv1.MPDCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mpdv1.MPDCluster), err
}

// Update takes the representation of a mPDCluster and updates it. Returns the server's representation of the mPDCluster, and an error, if there is any.
func (c *FakeMPDClusters) Update(mPDCluster *mpdv1.MPDCluster) (result *mpdv1.MPDCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mpdclustersResource, c.ns, mPDCluster), &mpdv1.MPDCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mpdv1.MPDCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMPDClusters) UpdateStatus(mPDCluster *mpdv1.MPDCluster) (*mpdv1.MPDCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mpdclustersResource, "status", c.ns, mPDCluster), &mpdv1.MPDCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mpdv1.MPDCluster), err
}

// Delete takes name of the mPDCluster and deletes it. Returns an error if one occurs.
func (c *FakeMPDClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mpdclustersResource, c.ns, name), &mpdv1.MPDCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMPDClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mpdclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mpdv1.MPDClusterList{})
	return err
}

// Patch applies the patch and returns the patched mPDCluster.
func (c *FakeMPDClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mpdv1.MPDCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mpdclustersResource, c.ns, name, pt, data, subresources...), &mpdv1.MPDCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mpdv1.MPDCluster), err
}
