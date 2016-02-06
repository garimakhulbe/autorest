﻿// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0. See License.txt in the project root for license information.

using System.Collections.Generic;
using System.Linq;

using Microsoft.Rest.Generator.ClientModel;
using Microsoft.Rest.Generator.Go;
using Microsoft.Rest.Generator.Utilities;

namespace Microsoft.Rest.Generator.Go
{
    public class ServiceClientTemplateModel : ServiceClient
    {
        public readonly string BaseClient;
        public string ClientName { get; set; }
        public string ClientDocumentation { get; set; }
        public readonly string MethodGroupName;
        public readonly string PackageName;

        public readonly MethodScopeProvider MethodScope;
        public readonly List<MethodTemplateModel> MethodTemplateModels;

        public ServiceClientTemplateModel(ServiceClient serviceClient, string packageName, string methodGroupName = null)
        {
            this.LoadFrom(serviceClient);

            MethodGroupName = methodGroupName == null
                                ? string.Empty
                                : methodGroupName;
            PackageName = packageName == null
                            ? string.Empty
                            : packageName;

            // TODO (gosdk): Derive client name from Swagger once the Swagger settles down.
            //BaseClient = serviceClient.Name.TrimPackageName(PackageName);
            //if (BaseClient == "Client")
            //{
            //    BaseClient = "ManagementClient";
            //}
            BaseClient = "ManagementClient";
            ClientName = string.IsNullOrEmpty(MethodGroupName)
                            ? BaseClient
                            : (MethodGroupName + "Client").TrimPackageName(PackageName);

            MethodScope = new MethodScopeProvider();
            MethodTemplateModels = new List<MethodTemplateModel>();
            Methods.Where(m => m.BelongsToGroup(MethodGroupName, PackageName))
                .OrderBy(m => m.Name)
                .ForEach(m => MethodTemplateModels.Add(new MethodTemplateModel(m, ClientName, PackageName, new MethodScopeProvider())));

            Documentation = string.Format("Package {0} implements the Azure ARM {1} service API version {2}.\n\n{3}", PackageName, ServiceName, ApiVersion,
                                    !string.IsNullOrEmpty(Documentation) ? Documentation.UnwrapAnchorTags() + "." : "");
            ClientDocumentation = string.Format("{0} is the base client for {1}.", ClientName, ServiceName);
        }

        public string ServiceName
        {
            get
            {
                if (!string.IsNullOrEmpty(PackageName))
                {
                    return GoCodeNamer.PascalCase(PackageName);
                }
                return string.Empty;
            }
        }

        public virtual IEnumerable<string> Imports
        {
            get
            {
                var imports = new HashSet<string>();
                imports.UnionWith(GoCodeNamer.AutorestImports);
                if (MethodTemplateModels.Count() > 0)
                {
                    imports.UnionWith(GoCodeNamer.StandardImports);
                    MethodTemplateModels
                        .ForEach(mtm =>
                        {
                            mtm.Parameters.ForEach(p => p.AddImports(imports));
                            if (mtm.HasReturnValue() && mtm.ReturnValue().Body != PrimaryType.Stream)
                            {
                                mtm.ReturnType.Body.AddImports(imports);
                            }
                        });
                }
                return imports.OrderBy(i => i);
            }
        }
    }
}