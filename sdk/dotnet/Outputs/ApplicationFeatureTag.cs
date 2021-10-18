// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class ApplicationFeatureTag
    {
        /// <summary>
        /// Whether this application represents a custom SAML application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag. Defaults to `false`.
        /// </summary>
        public readonly bool? CustomSingleSignOn;
        /// <summary>
        /// Whether this application represents an Enterprise Application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag. Defaults to `false`.
        /// </summary>
        public readonly bool? Enterprise;
        /// <summary>
        /// Whether this application represents a gallery application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag. Defaults to `false`.
        /// </summary>
        public readonly bool? Gallery;
        /// <summary>
        /// Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag. Defaults to `false`.
        /// </summary>
        public readonly bool? Hide;

        [OutputConstructor]
        private ApplicationFeatureTag(
            bool? customSingleSignOn,

            bool? enterprise,

            bool? gallery,

            bool? hide)
        {
            CustomSingleSignOn = customSingleSignOn;
            Enterprise = enterprise;
            Gallery = gallery;
            Hide = hide;
        }
    }
}
