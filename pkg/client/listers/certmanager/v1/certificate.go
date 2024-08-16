/*
Copyright The cert-manager Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// CertificateLister helps list Certificates.
// All objects returned here must be treated as read-only.
type CertificateLister interface {
	// List lists all Certificates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Certificate, err error)
	// Certificates returns an object that can list and get Certificates.
	Certificates(namespace string) CertificateNamespaceLister
	CertificateListerExpansion
}

// certificateLister implements the CertificateLister interface.
type certificateLister struct {
	listers.ResourceIndexer[*v1.Certificate]
}

// NewCertificateLister returns a new CertificateLister.
func NewCertificateLister(indexer cache.Indexer) CertificateLister {
	return &certificateLister{listers.New[*v1.Certificate](indexer, v1.Resource("certificate"))}
}

// Certificates returns an object that can list and get Certificates.
func (s *certificateLister) Certificates(namespace string) CertificateNamespaceLister {
	return certificateNamespaceLister{listers.NewNamespaced[*v1.Certificate](s.ResourceIndexer, namespace)}
}

// CertificateNamespaceLister helps list and get Certificates.
// All objects returned here must be treated as read-only.
type CertificateNamespaceLister interface {
	// List lists all Certificates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Certificate, err error)
	// Get retrieves the Certificate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Certificate, error)
	CertificateNamespaceListerExpansion
}

// certificateNamespaceLister implements the CertificateNamespaceLister
// interface.
type certificateNamespaceLister struct {
	listers.ResourceIndexer[*v1.Certificate]
}
