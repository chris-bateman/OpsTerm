// aws/session.go
package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// AuthMethod is the type of AWS authentication selected by the user.
type AuthMethod int

const (
	DefaultProfile AuthMethod = iota
	NamedProfile
	AssumeRole
	EnvVars
)

// AuthInput contains any user-entered values needed for config loading.
type AuthInput struct {
	Method      AuthMethod
	ProfileName string // for NamedProfile
	RoleArn     string // for AssumeRole
	Region      string
}

// LoadAWSConfig returns an AWS config based on the selected auth method.
func LoadAWSConfig(ctx context.Context, input AuthInput) (aws.Config, error) {
	var cfg aws.Config
	var err error

	switch input.Method {
	case DefaultProfile:
		cfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(input.Region))

	case NamedProfile:
		cfg, err = config.LoadDefaultConfig(ctx,
			config.WithSharedConfigProfile(input.ProfileName),
			config.WithRegion(input.Region),
		)

	case EnvVars:
		// No-op — SDK reads directly from env vars
		cfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(input.Region))

	case AssumeRole:
		// First load base config, then assume role
		baseCfg, baseErr := config.LoadDefaultConfig(ctx, config.WithRegion(input.Region))
		if baseErr != nil {
			return aws.Config{}, baseErr
		}

		stsClient := sts.NewFromConfig(baseCfg)
		creds := aws.NewCredentialsCache(stscreds.NewAssumeRoleProvider(stsClient, input.RoleArn))

		cfg = baseCfg
		cfg.Credentials = creds

	default:
		err = fmt.Errorf("unsupported auth method")
	}

	return cfg, err
}