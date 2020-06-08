//
// Copyright (c) 2019-2020 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

// The package that is used by components to get configuration.
//
// Typically each configuration property has the default value.
// Default value is supposed to be overridden via config map.
//
// There is the following configuration names convention:
// - words are lower-cased
// - . is used to separate subcomponents
// - _ is used to separate words in the component name
//
// Examples:
// che.workspace.plugin_broker.artifacts.image
// che.workspace.plugin_broker.artifacts.memory_limit
// Where:
// che.workspace - indicates that this is going to land to workspace runtime
// plugin_broker - is a part of workspace
// artifacts - is a type of plugin broker component
// memory_limit, image - are just different configuration properties for the same component
package config