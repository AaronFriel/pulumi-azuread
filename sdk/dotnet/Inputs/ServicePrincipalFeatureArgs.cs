// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ServicePrincipalFeatureArgs : Pulumi.ResourceArgs
    {
        [Input("customSingleSignOnApp")]
        public Input<bool>? CustomSingleSignOnApp { get; set; }

        [Input("enterpriseApplication")]
        public Input<bool>? EnterpriseApplication { get; set; }

        [Input("galleryApplication")]
        public Input<bool>? GalleryApplication { get; set; }

        [Input("visibleToUsers")]
        public Input<bool>? VisibleToUsers { get; set; }

        public ServicePrincipalFeatureArgs()
        {
        }
    }
}
