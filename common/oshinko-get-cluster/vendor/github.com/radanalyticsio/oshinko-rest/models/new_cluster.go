package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*NewCluster new cluster

swagger:model NewCluster
*/
type NewCluster struct {

	/* config
	 */
	Config *NewClusterConfig `json:"config,omitempty"`

	/* Unique name for the cluster

	Required: true
	*/
	Name *string `json:"name"`
}

// Validate validates this new cluster
func (m *NewCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewCluster) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {

		if err := m.Config.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *NewCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

/*NewClusterConfig Cluster configuration values

swagger:model NewClusterConfig
*/
type NewClusterConfig struct {

	/* The count of master nodes requested in the cluster (must be > 0)
	 */
	MasterCount int64 `json:"masterCount,omitempty"`

	/* The name of a stored cluster configuration
	 */
	Name interface{} `json:"name,omitempty"`

	/* The count of worker nodes requested in the cluster (must be > 0)
	 */
	WorkerCount int64 `json:"workerCount,omitempty"`
}

// Validate validates this new cluster config
func (m *NewClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
