// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	url "net/url"
	unsafe "unsafe"

	pedia "github.com/clusterpedia-io/clusterpedia/pkg/apis/pedia"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CollectionResource)(nil), (*pedia.CollectionResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CollectionResource_To_pedia_CollectionResource(a.(*CollectionResource), b.(*pedia.CollectionResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pedia.CollectionResource)(nil), (*CollectionResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pedia_CollectionResource_To_v1alpha1_CollectionResource(a.(*pedia.CollectionResource), b.(*CollectionResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CollectionResourceList)(nil), (*pedia.CollectionResourceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CollectionResourceList_To_pedia_CollectionResourceList(a.(*CollectionResourceList), b.(*pedia.CollectionResourceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pedia.CollectionResourceList)(nil), (*CollectionResourceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pedia_CollectionResourceList_To_v1alpha1_CollectionResourceList(a.(*pedia.CollectionResourceList), b.(*CollectionResourceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CollectionResourceType)(nil), (*pedia.CollectionResourceType)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CollectionResourceType_To_pedia_CollectionResourceType(a.(*CollectionResourceType), b.(*pedia.CollectionResourceType), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pedia.CollectionResourceType)(nil), (*CollectionResourceType)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pedia_CollectionResourceType_To_v1alpha1_CollectionResourceType(a.(*pedia.CollectionResourceType), b.(*CollectionResourceType), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*ListOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_url_Values_To_v1alpha1_ListOptions(a.(*url.Values), b.(*ListOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*pedia.ListOptions)(nil), (*ListOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pedia_ListOptions_To_v1alpha1_ListOptions(a.(*pedia.ListOptions), b.(*ListOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*url.Values)(nil), (*ListOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_url_Values_To_v1alpha1_ListOptions(a.(*url.Values), b.(*ListOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ListOptions)(nil), (*pedia.ListOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ListOptions_To_pedia_ListOptions(a.(*ListOptions), b.(*pedia.ListOptions), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CollectionResource_To_pedia_CollectionResource(in *CollectionResource, out *pedia.CollectionResource, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.ResourceTypes = *(*[]pedia.CollectionResourceType)(unsafe.Pointer(&in.ResourceTypes))
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]runtime.Object, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_CollectionResource_To_pedia_CollectionResource is an autogenerated conversion function.
func Convert_v1alpha1_CollectionResource_To_pedia_CollectionResource(in *CollectionResource, out *pedia.CollectionResource, s conversion.Scope) error {
	return autoConvert_v1alpha1_CollectionResource_To_pedia_CollectionResource(in, out, s)
}

func autoConvert_pedia_CollectionResource_To_v1alpha1_CollectionResource(in *pedia.CollectionResource, out *CollectionResource, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.ResourceTypes = *(*[]CollectionResourceType)(unsafe.Pointer(&in.ResourceTypes))
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_pedia_CollectionResource_To_v1alpha1_CollectionResource is an autogenerated conversion function.
func Convert_pedia_CollectionResource_To_v1alpha1_CollectionResource(in *pedia.CollectionResource, out *CollectionResource, s conversion.Scope) error {
	return autoConvert_pedia_CollectionResource_To_v1alpha1_CollectionResource(in, out, s)
}

func autoConvert_v1alpha1_CollectionResourceList_To_pedia_CollectionResourceList(in *CollectionResourceList, out *pedia.CollectionResourceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]pedia.CollectionResource, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_CollectionResource_To_pedia_CollectionResource(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_CollectionResourceList_To_pedia_CollectionResourceList is an autogenerated conversion function.
func Convert_v1alpha1_CollectionResourceList_To_pedia_CollectionResourceList(in *CollectionResourceList, out *pedia.CollectionResourceList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CollectionResourceList_To_pedia_CollectionResourceList(in, out, s)
}

func autoConvert_pedia_CollectionResourceList_To_v1alpha1_CollectionResourceList(in *pedia.CollectionResourceList, out *CollectionResourceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CollectionResource, len(*in))
		for i := range *in {
			if err := Convert_pedia_CollectionResource_To_v1alpha1_CollectionResource(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_pedia_CollectionResourceList_To_v1alpha1_CollectionResourceList is an autogenerated conversion function.
func Convert_pedia_CollectionResourceList_To_v1alpha1_CollectionResourceList(in *pedia.CollectionResourceList, out *CollectionResourceList, s conversion.Scope) error {
	return autoConvert_pedia_CollectionResourceList_To_v1alpha1_CollectionResourceList(in, out, s)
}

func autoConvert_v1alpha1_CollectionResourceType_To_pedia_CollectionResourceType(in *CollectionResourceType, out *pedia.CollectionResourceType, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	out.Kind = in.Kind
	out.Resource = in.Resource
	return nil
}

// Convert_v1alpha1_CollectionResourceType_To_pedia_CollectionResourceType is an autogenerated conversion function.
func Convert_v1alpha1_CollectionResourceType_To_pedia_CollectionResourceType(in *CollectionResourceType, out *pedia.CollectionResourceType, s conversion.Scope) error {
	return autoConvert_v1alpha1_CollectionResourceType_To_pedia_CollectionResourceType(in, out, s)
}

func autoConvert_pedia_CollectionResourceType_To_v1alpha1_CollectionResourceType(in *pedia.CollectionResourceType, out *CollectionResourceType, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	out.Kind = in.Kind
	out.Resource = in.Resource
	return nil
}

// Convert_pedia_CollectionResourceType_To_v1alpha1_CollectionResourceType is an autogenerated conversion function.
func Convert_pedia_CollectionResourceType_To_v1alpha1_CollectionResourceType(in *pedia.CollectionResourceType, out *CollectionResourceType, s conversion.Scope) error {
	return autoConvert_pedia_CollectionResourceType_To_v1alpha1_CollectionResourceType(in, out, s)
}

func autoConvert_v1alpha1_ListOptions_To_pedia_ListOptions(in *ListOptions, out *pedia.ListOptions, s conversion.Scope) error {
	// FIXME: Provide conversion function to convert v1.ListOptions to internalversion.ListOptions
	compileErrorOnMissingConversion()
	// WARNING: in.Names requires manual conversion: inconvertible types (string vs []string)
	out.Owner = in.Owner
	// WARNING: in.ClusterNames requires manual conversion: inconvertible types (string vs []string)
	// WARNING: in.Namespaces requires manual conversion: inconvertible types (string vs []string)
	// WARNING: in.OrderBy requires manual conversion: inconvertible types (string vs []github.com/clusterpedia-io/clusterpedia/pkg/apis/pedia.OrderBy)
	return nil
}

func autoConvert_pedia_ListOptions_To_v1alpha1_ListOptions(in *pedia.ListOptions, out *ListOptions, s conversion.Scope) error {
	// FIXME: Provide conversion function to convert internalversion.ListOptions to v1.ListOptions
	compileErrorOnMissingConversion()
	if err := runtime.Convert_Slice_string_To_string(&in.Names, &out.Names, s); err != nil {
		return err
	}
	out.Owner = in.Owner
	if err := runtime.Convert_Slice_string_To_string(&in.ClusterNames, &out.ClusterNames, s); err != nil {
		return err
	}
	if err := runtime.Convert_Slice_string_To_string(&in.Namespaces, &out.Namespaces, s); err != nil {
		return err
	}
	// WARNING: in.OrderBy requires manual conversion: inconvertible types ([]github.com/clusterpedia-io/clusterpedia/pkg/apis/pedia.OrderBy vs string)
	// WARNING: in.ExtraLabelSelector requires manual conversion: does not exist in peer-type
	// WARNING: in.ExtraQuery requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_url_Values_To_v1alpha1_ListOptions(in *url.Values, out *ListOptions, s conversion.Scope) error {
	// WARNING: Field ListOptions does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["names"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Names, s); err != nil {
			return err
		}
	} else {
		out.Names = ""
	}
	if values, ok := map[string][]string(*in)["owner"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Owner, s); err != nil {
			return err
		}
	} else {
		out.Owner = ""
	}
	if values, ok := map[string][]string(*in)["clusters"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.ClusterNames, s); err != nil {
			return err
		}
	} else {
		out.ClusterNames = ""
	}
	if values, ok := map[string][]string(*in)["namespaces"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Namespaces, s); err != nil {
			return err
		}
	} else {
		out.Namespaces = ""
	}
	if values, ok := map[string][]string(*in)["orderby"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.OrderBy, s); err != nil {
			return err
		}
	} else {
		out.OrderBy = ""
	}
	return nil
}
