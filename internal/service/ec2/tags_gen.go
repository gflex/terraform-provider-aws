// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
)

// GetTag fetches an individual ec2 service tag for a resource.
// Returns whether the key value and any errors. A NotFoundError is used to signal that no value was found.
// This function will optimise the handling over ListTags, if possible.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func GetTag(ctx context.Context, conn ec2iface.EC2API, identifier, key string) (*string, error) {
	input := &ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("resource-id"),
				Values: []*string{aws.String(identifier)},
			},
			{
				Name:   aws.String("key"),
				Values: []*string{aws.String(key)},
			},
		},
	}

	output, err := conn.DescribeTagsWithContext(ctx, input)

	if err != nil {
		return nil, err
	}

	listTags := KeyValueTags(ctx, output.Tags)

	if !listTags.KeyExists(key) {
		return nil, tfresource.NewEmptyResultError(nil)
	}

	return listTags.KeyValue(key), nil
}

// ListTags lists ec2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn ec2iface.EC2API, identifier string) (tftags.KeyValueTags, error) {
	input := &ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("resource-id"),
				Values: []*string{aws.String(identifier)},
			},
		},
	}

	output, err := conn.DescribeTagsWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// ListTags lists ec2 service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := ListTags(ctx, meta.(*conns.AWSClient).EC2Conn(), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns ec2 service tags.
func Tags(tags tftags.KeyValueTags) []*ec2.Tag {
	result := make([]*ec2.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &ec2.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from ec2 service tags.
//
// Accepts the following types:
//   - []*ec2.Tag
//   - []*ec2.TagDescription
func KeyValueTags(ctx context.Context, tags any) tftags.KeyValueTags {
	switch tags := tags.(type) {
	case []*ec2.Tag:
		m := make(map[string]*string, len(tags))

		for _, tag := range tags {
			m[aws.StringValue(tag.Key)] = tag.Value
		}

		return tftags.New(ctx, m)
	case []*ec2.TagDescription:
		m := make(map[string]*string, len(tags))

		for _, tag := range tags {
			m[aws.StringValue(tag.Key)] = tag.Value
		}

		return tftags.New(ctx, m)
	default:
		return tftags.New(ctx, nil)
	}
}

// GetTagsIn returns ec2 service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*ec2.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets ec2 service tags in Context.
func SetTagsOut(ctx context.Context, tags any) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// UpdateTags updates ec2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.

func UpdateTags(ctx context.Context, conn ec2iface.EC2API, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &ec2.DeleteTagsInput{
			Resources: aws.StringSlice([]string{identifier}),
			Tags:      Tags(removedTags.IgnoreAWS()),
		}

		_, err := conn.DeleteTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &ec2.CreateTagsInput{
			Resources: aws.StringSlice([]string{identifier}),
			Tags:      Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.CreateTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates ec2 service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).EC2Conn(), identifier, oldTags, newTags)
}
