// Copyright (c) 2022 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"testing"

	"github.com/alibaba/higress/test/ingress/conformance/utils/http"
	"github.com/alibaba/higress/test/ingress/conformance/utils/roundtripper"
	"github.com/alibaba/higress/test/ingress/conformance/utils/suite"
)

func init() {
	HigressConformanceTests = append(HigressConformanceTests, HttpRedirectAsHttps)
}

var HttpRedirectAsHttps = suite.ConformanceTest{
	ShortName:   "HttpRedirectAsHttps",
	Description: "The Ingress in the higress-conformance-infra namespace Server-side HTTPS enforcement through redirect.",
	Manifests:   []string{"tests/httproute-redirct-as-https.yaml"},
	Test: func(t *testing.T, suite *suite.ConformanceTestSuite) {
		testcases := []http.Assertion{
			{
				Meta: http.AssertionMeta{
					TargetBackend:   "infra-backend-v1",
					TargetNamespace: "higress-conformance-infra",
				},
				Request: http.AssertionRequest{
					ActualRequest: http.Request{
						Host:             "test.com",
						Path:             "/test",
						UnfollowRedirect: true,
					},
					RedirectRequest: &roundtripper.RedirectRequest{
						Scheme: "https",
						Host:   "test.com",
						Path:   "/test",
					},
				},
				Response: http.AssertionResponse{
					ExpectedResponse: http.Response{
						StatusCode: 308,
					},
				},
			},
		}

		t.Run("HttpRedirectAsHttps", func(t *testing.T) {
			for _, testcase := range testcases {
				http.MakeRequestAndExpectEventuallyConsistentResponse(t, suite.RoundTripper, suite.TimeoutConfig, suite.GatewayAddress, testcase)
			}
		})

	},
}
