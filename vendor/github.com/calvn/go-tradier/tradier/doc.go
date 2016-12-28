// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

/*
Package tradier provides a client for interacting with the Tradier API.

Usage:

    import "github.com/google/go-github/github"

Getting started

The quickest way to get started is to request a non-expiring access token (or more formally, a refresh token) from Tradier. Initantiate a new Tradier client and provide an http.Client that has the proper access token. Use the different services to call specific endpoints.

    import "golang.org/x/oauth2"

    func main() {
      ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: "... your access token ..."},
      )
      tc := oauth2.NewClient(oauth2.NoContext, ts)

      client := tradier.NewClient(tc)

      // Returns the profile of the user
      profile, _, err := client.User.Profile()
    }

If the application is frontend-facing, then the authorization grant flow (https://tools.ietf.org/html/rfc6749#section-1.3.1) should be used to request an access token from the resource owner (the user). Detailed information regarding this process can be found on https://developer.tradier.com/documentation/oauth/getting-started.

Resources and structs

All resources from Tradier are mapped to its corresponding structs. These structs use pointer values in their fields except in the case of slices, which is referenced directly and hold a collection of pointer values. This is to ensure that there is a difference between a nil and zero value when marshalling a resource into its JSON representation.

In most cases (except for the /profile endpoint), the Tradier API returns a collection of objects differently, depending on whether the response is an array of objects, a single object, or an empty collection. For example:

    // Array of objects
    {
      orders: {
        order: [
          ...
        ]
      }
    }

    // Single object
    {
      orders: {
        order: {
          ...
        }
      }
    }

    // Empty collection
    {
      orders: "null"
    }

There are structs, such as Orders, Watchlists, and Accounts, that represent these collections. In cases where a service method returns these structs, the underlying object can be accessed via indexing or length-checking. For example, if orders is an instance of Orders, acessing each of the orders can be done like so:

  - If order is an array of objects, it can be accessed though indexing on orders[i].
  - If order is a single object, it can be accessed on the zeroth element (i.e. orders[0]).
  - If order is "null", orders length will be zero, len(orders) == 0.

*/
package tradier
