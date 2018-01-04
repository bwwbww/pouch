// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContainerJSON ContainerJSON contains response of Engine API:
// GET "/containers/{id}/json"
//
// swagger:model ContainerJSON

type ContainerJSON struct {

	// app armor profile
	AppArmorProfile string `json:"AppArmorProfile,omitempty"`

	// The arguments to the command being run
	Args []string `json:"Args"`

	// config
	Config *ContainerConfig `json:"Config,omitempty"`

	// The time the container was created
	Created string `json:"Created,omitempty"`

	// driver
	Driver string `json:"Driver,omitempty"`

	// exec ids
	ExecIds string `json:"ExecIDs,omitempty"`

	// graph driver
	GraphDriver *GraphDriverData `json:"GraphDriver,omitempty"`

	// host config
	HostConfig *HostConfig `json:"HostConfig,omitempty"`

	// hostname path
	HostnamePath string `json:"HostnamePath,omitempty"`

	// hosts path
	HostsPath string `json:"HostsPath,omitempty"`

	// The ID of the container
	ID string `json:"Id,omitempty"`

	// The container's image
	Image string `json:"Image,omitempty"`

	// log path
	LogPath string `json:"LogPath,omitempty"`

	// mount label
	MountLabel string `json:"MountLabel,omitempty"`

	// Set of mount point in a container.
	Mounts []MountPoint `json:"Mounts"`

	// name
	Name string `json:"Name,omitempty"`

	// NetworkSettings exposes the network settings in the API.
	NetworkSettings *NetworkSettings `json:"NetworkSettings,omitempty"`

	// The path to the command being run
	Path string `json:"Path,omitempty"`

	// process label
	ProcessLabel string `json:"ProcessLabel,omitempty"`

	// resolv conf path
	ResolvConfPath string `json:"ResolvConfPath,omitempty"`

	// restart count
	RestartCount int64 `json:"RestartCount,omitempty"`

	// The total size of all the files in this container.
	SizeRootFs *int64 `json:"SizeRootFs,omitempty"`

	// The size of files that have been created or changed by this container.
	SizeRw *int64 `json:"SizeRw,omitempty"`

	// The state of the container.
	State *ContainerState `json:"State,omitempty"`
}

/* polymorph ContainerJSON AppArmorProfile false */

/* polymorph ContainerJSON Args false */

/* polymorph ContainerJSON Config false */

/* polymorph ContainerJSON Created false */

/* polymorph ContainerJSON Driver false */

/* polymorph ContainerJSON ExecIDs false */

/* polymorph ContainerJSON GraphDriver false */

/* polymorph ContainerJSON HostConfig false */

/* polymorph ContainerJSON HostnamePath false */

/* polymorph ContainerJSON HostsPath false */

/* polymorph ContainerJSON Id false */

/* polymorph ContainerJSON Image false */

/* polymorph ContainerJSON LogPath false */

/* polymorph ContainerJSON MountLabel false */

/* polymorph ContainerJSON Mounts false */

/* polymorph ContainerJSON Name false */

/* polymorph ContainerJSON NetworkSettings false */

/* polymorph ContainerJSON Path false */

/* polymorph ContainerJSON ProcessLabel false */

/* polymorph ContainerJSON ResolvConfPath false */

/* polymorph ContainerJSON RestartCount false */

/* polymorph ContainerJSON SizeRootFs false */

/* polymorph ContainerJSON SizeRw false */

/* polymorph ContainerJSON State false */

// Validate validates this container JSON
func (m *ContainerJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGraphDriver(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkSettings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerJSON) validateArgs(formats strfmt.Registry) error {

	if swag.IsZero(m.Args) { // not required
		return nil
	}

	return nil
}

func (m *ContainerJSON) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {

		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerJSON) validateGraphDriver(formats strfmt.Registry) error {

	if swag.IsZero(m.GraphDriver) { // not required
		return nil
	}

	if m.GraphDriver != nil {

		if err := m.GraphDriver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GraphDriver")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerJSON) validateMounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Mounts) { // not required
		return nil
	}

	return nil
}

func (m *ContainerJSON) validateNetworkSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkSettings) { // not required
		return nil
	}

	if m.NetworkSettings != nil {

		if err := m.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkSettings")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerJSON) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {

		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("State")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerJSON) UnmarshalBinary(b []byte) error {
	var res ContainerJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
