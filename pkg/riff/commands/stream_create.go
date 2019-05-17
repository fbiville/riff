/*
 * Copyright 2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
	"context"
	"github.com/projectriff/riff/pkg/validation"
	streamv1alpha1 "github.com/projectriff/system/pkg/apis/stream/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/projectriff/riff/pkg/cli"
	"github.com/spf13/cobra"
)

type StreamCreateOptions struct {
	ContentType string // TODO update Stream CRD
	Name        string
	Namespace   string
	Provider    string
}

func (opts *StreamCreateOptions) Validate(ctx context.Context) *cli.FieldError {
	errs := &cli.FieldError{}
	if opts.Namespace == "" {
		errs = errs.Also(cli.ErrMissingField(cli.NamespaceFlagName))
	}
	if opts.Name == "" {
		errs = errs.Also(cli.ErrMissingField(cli.NameFlagName))
	} else {
		errs = errs.Also(validation.K8sName(opts.Name, cli.NameFlagName))
	}
	if opts.Provider == "" {
		errs = errs.Also(cli.ErrMissingField(cli.ProviderFlagName))
	}
	return errs
}

func (opts *StreamCreateOptions) Exec(ctx context.Context, c *cli.Config) error {
	stream := &streamv1alpha1.Stream{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: opts.Namespace,
			Name:      opts.Name,
		},
		Spec: streamv1alpha1.StreamSpec{
			Provider: opts.Provider,
		},
	}
	stream, err := c.Stream().Streams(opts.Namespace).Create(stream)
	if err != nil {
		return err
	}

	_, _ = c.Successf("Created stream %q\n", stream.Name)
	return nil
}

func NewStreamCreateCommand(c *cli.Config) *cobra.Command {
	opts := &StreamCreateOptions{}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   "<todo>",
		Example: "<todo>",
		Args:    cli.Args(),
		PreRunE: cli.ValidateOptions(opts),
		RunE:    cli.ExecOptions(c, opts),
	}

	cli.NamespaceFlag(cmd, c, &opts.Namespace)

	return cmd
}
