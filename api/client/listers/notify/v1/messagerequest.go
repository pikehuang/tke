/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "tkestack.io/tke/api/notify/v1"
)

// MessageRequestLister helps list MessageRequests.
// All objects returned here must be treated as read-only.
type MessageRequestLister interface {
	// List lists all MessageRequests in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.MessageRequest, err error)
	// MessageRequests returns an object that can list and get MessageRequests.
	MessageRequests(namespace string) MessageRequestNamespaceLister
	MessageRequestListerExpansion
}

// messageRequestLister implements the MessageRequestLister interface.
type messageRequestLister struct {
	indexer cache.Indexer
}

// NewMessageRequestLister returns a new MessageRequestLister.
func NewMessageRequestLister(indexer cache.Indexer) MessageRequestLister {
	return &messageRequestLister{indexer: indexer}
}

// List lists all MessageRequests in the indexer.
func (s *messageRequestLister) List(selector labels.Selector) (ret []*v1.MessageRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MessageRequest))
	})
	return ret, err
}

// MessageRequests returns an object that can list and get MessageRequests.
func (s *messageRequestLister) MessageRequests(namespace string) MessageRequestNamespaceLister {
	return messageRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MessageRequestNamespaceLister helps list and get MessageRequests.
// All objects returned here must be treated as read-only.
type MessageRequestNamespaceLister interface {
	// List lists all MessageRequests in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.MessageRequest, err error)
	// Get retrieves the MessageRequest from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.MessageRequest, error)
	MessageRequestNamespaceListerExpansion
}

// messageRequestNamespaceLister implements the MessageRequestNamespaceLister
// interface.
type messageRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MessageRequests in the indexer for a given namespace.
func (s messageRequestNamespaceLister) List(selector labels.Selector) (ret []*v1.MessageRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MessageRequest))
	})
	return ret, err
}

// Get retrieves the MessageRequest from the indexer for a given namespace and name.
func (s messageRequestNamespaceLister) Get(name string) (*v1.MessageRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("messagerequest"), name)
	}
	return obj.(*v1.MessageRequest), nil
}
