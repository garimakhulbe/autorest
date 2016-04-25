﻿// Copyright (c) Microsoft Open Technologies, Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0. See License.txt in the project root for license information.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using Microsoft.Rest.Generator.Azure;
using Microsoft.Rest.Generator.ClientModel;
using Microsoft.Rest.Generator.Utilities;

namespace Microsoft.Rest.Generator.Go
{
    public static class Extensions
    {
        public const string ApiVersionName = "APIVersion";
        public const string ApiVersionSerializedName = "api-version";

        public const string SkipUrlEncoding = "x-ms-skip-url-encoding";

        private static readonly Regex SplitPattern = new Regex(@"(\p{Lu}\p{Ll}+)");

        /////////////////////////////////////////////////////////////////////////////////////////
        //
        // General Extensions
        //
        /////////////////////////////////////////////////////////////////////////////////////////

        public static string ToSentence(this string value)
        {
            if (string.IsNullOrWhiteSpace(value))
            {
                return string.Empty;
            }
            else
            {
                value = value.Trim();
                return value.First().ToString().ToLowerInvariant() + (value.Length > 1 ? value.Substring(1) : "");
            }
        }

        public static string Capitalize(this string value)
        {
            return string.IsNullOrWhiteSpace(value)
                    ? string.Empty
                    : value.First()
                           .ToString()
                           .ToUpperInvariant() + (value.Length > 1 ? value.Substring(1) : "");
        }

        public static string ToPhrase(this string value)
        {
            List<string> words = new List<string>(value.ToWords());
            for (int i = 0; i < words.Count; i++)
            {
                words[i] = words[i].ToLowerInvariant();
            }
            return string.Join(" ", words.ToArray());
        }
        
        public static string[] ToWords(this string value)
        {
            return SplitPattern.Split(value).Where(s => !string.IsNullOrEmpty(s)).ToArray();
        }

        public static string TrimStartsWith(this string value, string s)
        {
            if (!string.IsNullOrEmpty(s) && s.Length < value.Length && value.StartsWith(s, StringComparison.InvariantCultureIgnoreCase))
            {
                value = value.Substring(s.Length);
            }
            return value;
        }

        public static string TrimPackageName(this string value, string packageName)
        {
            var cchValue = value.Length;
            value = value.TrimStartsWith(packageName);
            if (value.Length == cchValue && packageName.EndsWith("s") && (value.Length - packageName.Length + 1) > 1)
            {
                value = value.TrimStartsWith(packageName.Substring(0, packageName.Length-1));
            }
            return value;
        }

        public static string EmitAsArguments(this IList<string> arguments)
        {
            var sb = new StringBuilder();

            if (arguments.Count() > 0)
            {
                sb.Append(arguments[0]);

                for (var i = 1; i < arguments.Count(); i++)
                {
                    sb.AppendLine(",");
                    sb.Append(arguments[i]);
                }
            }

            return sb.ToString();
        }


        // This function removes html anchor tags and reformats the comment text.
        // For example, Swagger documentation text --> "This is a documentation text. For information see <a href=LINK">CONTENT.</a>"
        // reformats to  --> "This is a documentation text. For information see CONTENT (LINK)."
        public static string UnwrapAnchorTags(this string comments)
        {
            string pattern = "([^<>]*)<a\\s*.*\\shref\\s*=\\s*[\'\"]([^\'\"]*)[\'\"][^>]*>(.*)</a>";
            Regex r = new Regex(pattern);
            Match match = r.Match(comments);

            if (match.Success)
            {
                string content = match.Groups[3].Value;
                string link = match.Groups[2].Value;

                return (".?!;:".Contains(content[content.Length - 1])
                        ? match.Groups[1].Value + content.Substring(0, content.Length - 1) + " (" + link + ")" + content[content.Length - 1]
                        : match.Groups[1].Value + content + " (" + link + ")");
            }

            return comments;
        }
        
        /////////////////////////////////////////////////////////////////////////////////////////
        //
        // Parameter Extensions
        //
        /////////////////////////////////////////////////////////////////////////////////////////

        /// <summary>
        /// Return a Go map of required parameters.
        /// </summary>
        /// <param name="parameters"></param>
        /// <param name="mapVariable"></param>
        /// <returns></returns>
        public static string BuildParameterMap(this IEnumerable<Parameter> parameters, string mapVariable)
        {
            var builder = new StringBuilder();

            builder.Append(mapVariable);
            builder.Append(" := map[string]interface{} {");

            if (parameters.Count() > 0)
            {
                builder.AppendLine();
                var indented = new IndentedStringBuilder("  ");
                parameters
                    .Where(p => p.IsRequired)
                    .OrderBy(p => p.SerializedName)
                    .ForEach(p => indented.AppendLine("\"{0}\": {1},", p.NameForMap(), p.ValueForMap()));
                builder.Append(indented);
            }
            builder.AppendLine("}");
            return builder.ToString();
        }

        public static string AddToMap(this Parameter parameter, string mapVariable)
        {
            return string.Format("{0}[\"{1}\"] = {2}", mapVariable, parameter.NameForMap(), parameter.ValueForMap());
        }

        public static void AddImports(this Parameter parameter, HashSet<string> imports)
        {
            parameter.Type.AddImports(imports);
            /*if (parameter.RequiresUrlEncoding())
            {
                imports.Add("net/url");
            }*/
        }

        public static string NameForMap(this Parameter parameter)
        {
            return parameter.SerializedName.IsApiVersion()
                        ? ApiVersionSerializedName
                        : parameter.SerializedName;
        }

        public static string ValueForMap(this Parameter parameter)
        {
            if (parameter.SerializedName.IsApiVersion())
            {
                return "client."+ ApiVersionName;
            }
            var value = parameter.IsClientProperty()
                ? "client." + GoCodeNamer.PascalCase(parameter.Name)
                : parameter.Name;

            var format = parameter.IsRequired || parameter.Type.CanBeEmpty()
                                            ? "{0}"
                                            : "*{0}";

            var s = parameter.CollectionFormat != CollectionFormat.None
                                    ? string.Format("autorest.Stringify({0},\"{1}\")", format, parameter.CollectionFormat.GetSeparator()) 
                                    : string.Format("autorest.Stringify({0})", format);

            return string.Format(
                       parameter.RequiresUrlEncoding()
                                   ? string.Format("autorest.EncodeURI({0})",s)
                                   : string.Format("{0}", s), 
                       value);
        }

        public static IEnumerable<Parameter> ByLocation(this IEnumerable<Parameter> parameters, ParameterLocation location)
        {
            return parameters
                .Where(p => p.Location == location)
                .AsEnumerable();
        }

        public static IEnumerable<Parameter> ByLocationAsRequired(this IEnumerable<Parameter> parameters, ParameterLocation location, bool isRequired)
        {
            return parameters
                .Where(p => p.Location == location && p.IsRequired == isRequired)
                .AsEnumerable();
        }

        public static Parameter BodyParameter(this IEnumerable<Parameter> parameters)
        {
            var bodyParameters = parameters.ByLocation(ParameterLocation.Body);
            return bodyParameters.Count() > 0
                    ? bodyParameters.First()
                    : null;
        }

        public static IEnumerable<Parameter> BodyParameters(this IEnumerable<Parameter> parameters, bool isRequired)
        {
            return parameters.ByLocationAsRequired(ParameterLocation.Body, isRequired);
        }

        public static IEnumerable<Parameter> FormDataParameters(this IEnumerable<Parameter> parameters)
        {
            return parameters.ByLocation(ParameterLocation.FormData);
        }

        public static IEnumerable<Parameter> HeaderParameters(this IEnumerable<Parameter> parameters)
        {
            return parameters.ByLocation(ParameterLocation.Header);
        }

        public static IEnumerable<Parameter> PathParameters(this IEnumerable<Parameter> parameters)
        {
            return parameters.ByLocation(ParameterLocation.Path);
        }

        public static IEnumerable<Parameter> QueryParameters(this IEnumerable<Parameter> parameters)
        {
            return parameters.ByLocation(ParameterLocation.Query);
        }

        public static IEnumerable<Parameter> QueryParameters(this IEnumerable<Parameter> parameters, bool isRequired)
        {
            return parameters.ByLocationAsRequired(ParameterLocation.Query, isRequired);
        }

        /// <summary>
        /// Return the separator associated with a given collectionFormat
        /// </summary>
        /// <param name="format">The collection format</param>
        /// <returns>The separator</returns>
        private static string GetSeparator(this CollectionFormat format)
        {
            switch (format)
            {
                case CollectionFormat.Csv:
                    return ",";
                case CollectionFormat.Pipes:
                    return "|";
                case CollectionFormat.Ssv:
                    return " ";
                case CollectionFormat.Tsv:
                    return "\t";
                default:
                    throw new NotSupportedException(string.Format("Collection format {0} is not supported.", format));
            }
        }

        public static bool IsClientProperty(this Parameter parameter)
        {
            return parameter.ClientProperty != null;
        }

        public static bool IsMethodArgument(this Parameter parameter)
        {
            return !parameter.IsClientProperty();
        }

        public static bool IsApiVersion(this string name)
        {
            return name.Equals(ApiVersionSerializedName, StringComparison.OrdinalIgnoreCase);
        }

        public static bool RequiresUrlEncoding(this Parameter parameter)
        {
            return parameter.Location == ParameterLocation.Path && !parameter.Extensions.ContainsKey(SkipUrlEncoding);
        }

        public static string JsonTag(this Property property, bool omitEmpty = true)
        {
            return string.Format("`json:\"{0}{1}\"`", property.SerializedName, omitEmpty ? ",omitempty" : "");
        }

        /////////////////////////////////////////////////////////////////////////////////////////
        //
        // Type Extensions
        //
        /////////////////////////////////////////////////////////////////////////////////////////

        public static bool IsStreamType(this IType body)
        {
            var r = body as SyntheticType;
            return (r != null && (r.BaseType == PrimaryType.Stream))
                            ? true
                            : false;
        }

        public static bool CanBeEmpty(this IType type)
        {
            var dictionaryType = type as DictionaryType;
            var interfaceType = type as InterfaceType;
            var primaryType = type as PrimaryType;
            var sequenceType = type as SequenceType;
            var enumType = type as EnumType;

            return dictionaryType != null
                || interfaceType !=  null
                || (    primaryType != null
                    && (    primaryType == PrimaryType.ByteArray
                        ||  primaryType == PrimaryType.Stream
                        ||  primaryType == PrimaryType.String))
                || sequenceType != null
                || enumType != null;
        }

        public static bool CanBeNull(this IType type)
        {
            var dictionaryType = type as DictionaryType;
            var interfaceType = type as InterfaceType;
            var primaryType = type as PrimaryType;
            var sequenceType = type as SequenceType;

            return dictionaryType != null
                || interfaceType != null
                || (    primaryType != null
                    &&  (   primaryType == PrimaryType.ByteArray
                        ||  primaryType == PrimaryType.Stream))
                || sequenceType != null;
        }
        
        public static string GetEmptyCheck(this IType type, string valueReference, bool asEmpty = true)
        {
            if (type is PrimaryType)
            {
                return (type as PrimaryType).GetEmptyCheck(valueReference, asEmpty);
            }
            else if (type is SequenceType)
            {
                return (type as SequenceType).GetEmptyCheck(valueReference, asEmpty);
            }
            else if (type is DictionaryType)
            {
                return (type as DictionaryType).GetEmptyCheck(valueReference, asEmpty);
            }
            else if (type is EnumType)
            {
                return (type as EnumType).GetEmptyCheck(valueReference, asEmpty);
            }
            else
            {
                return string.Format(asEmpty
                                        ? "{0} == nil"
                                        : "{0} != nil", valueReference);
            }
        }

        public static string GetEmptyCheck(this DictionaryType type, string valueReference, bool asEmpty)
        {
            return string.Format(asEmpty
                                    ? "{0} == nil || len({0}) == 0"
                                    : "{0} != nil && len({0}) > 0", valueReference);
        }

        public static string GetEmptyCheck(this EnumType type, string valueReference, bool asEmpty)
        {
            return string.Format(asEmpty
                                    ? "len(string({0})) == 0"
                                    : "len(string({0})) > 0", valueReference);
        }

        public static string GetEmptyCheck(this PrimaryType type, string valueReference, bool asEmpty)
        {
            if (type == PrimaryType.ByteArray)
            {
                return string.Format(asEmpty
                                        ? "{0} == nil || len({0}) == 0"
                                        : "{0} != nil && len({0}) > 0", valueReference);
            }
            else if (type == PrimaryType.String)
            {
                return string.Format(asEmpty
                                        ? "len({0}) == 0"
                                        : "len({0}) > 0", valueReference);
            }
            else
            {
                return string.Format(asEmpty
                                        ? "{0} == nil"
                                        : "{0} != nil", valueReference);
            }
        }

        public static string GetEmptyCheck(this SequenceType type, string valueReference, bool asEmpty)
        {
            return string.Format(asEmpty
                                    ? "{0} == nil || len({0}) == 0"
                                    : "{0} != nil && len({0}) > 0", valueReference);
        }

        public static string Fields(this CompositeType compositeType)
        {
            var indented = new IndentedStringBuilder("    ");
            var properties = compositeType.Properties;

            if (compositeType.BaseModelType != null)
            {
                indented.Append(compositeType.BaseModelType.Fields());
            }

            // If the type is a paged model type, ensure the nextLink field exists
            // Note: Inject the field into a copy of the property list so as to not pollute the original list
            if (    compositeType is ModelTemplateModel
                &&  !String.IsNullOrEmpty((compositeType as ModelTemplateModel).NextLink))
            {
                var nextLinkField = (compositeType as ModelTemplateModel).NextLink;
                if (!properties.Any(p => p.Name == nextLinkField))
                {
                    var property = new Property();
                    property.Name = nextLinkField;
                    property.Type = PrimaryType.String;
                    properties = new List<Property>(properties);
                    properties.Add(property);
                }
            }

            // Emit each property, except for named Enumerated types, as a pointer to the type
            foreach (var property in properties)
            {
                EnumType enumType = property.Type as EnumType;
                if (enumType != null && enumType.IsNamed())
                {
                    indented.AppendFormat("{0} {1} {2}\n",
                                    property.Name,
                                    enumType.Name,
                                    property.JsonTag());

                }
                else if (property.Type is DictionaryType)
                {
                    indented.AppendFormat("{0} *{1} {2}\n", property.Name, (property.Type as MapType).FieldName, property.JsonTag());
                }
                else
                {
                    indented.AppendFormat("{0} *{1} {2}\n", property.Name, property.Type.Name, property.JsonTag());
                }
            }
            
            return indented.ToString();
        }

        public static string PreparerMethodName(this CompositeType compositeType)
        {
            return compositeType.Name + "Preparer";
        }

        public static bool IsNamed(this EnumType enumType)
        {
            return !string.IsNullOrEmpty(enumType.Name)
                && enumType.Name != "string"
                && enumType.Values.Count > 0;
        }

        public static void AddImports(this IType type, HashSet<string> imports)
        {
            if (type is DictionaryType)
            {
                (type as DictionaryType).AddImports(imports);
            }
            else if (type is PackageType)
            {
                (type as PackageType).AddImports(imports);
            }
            else if (type is PrimaryType)
            {
                (type as PrimaryType).AddImports(imports);
            }
            else if (type is SequenceType)
            {
                (type as SequenceType).AddImports(imports);
            }
        }

        public static void AddImports(this CompositeType compositeType, HashSet<string> imports)
        {
            compositeType
                .Properties
                .ForEach(p => p.Type.AddImports(imports));
        }

        public static void AddImports(this DictionaryType dictionaryType, HashSet<string> imports)
        {
            dictionaryType.ValueType.AddImports(imports);
        }

        public static void AddImports(this PackageType packageType, HashSet<string> imports)
        {
            imports.Add(packageType.Import);
        }

        public static void AddImports(this PrimaryType primaryType, HashSet<string> imports)
        {
            if (primaryType == PrimaryType.Stream)
            {
                imports.Add("io");
            }
        }

        public static void AddImports(this SequenceType sequenceType, HashSet<string> imports)
        {
            sequenceType.ElementType.AddImports(imports);
        }

        /// <summary>
        /// Generate code to perform required validation on a type
        /// </summary>
        /// <param name="type">The type to validate</param>
        /// <param name="scope">A scope provider for generating variable names as necessary</param>
        /// <param name="valueReference">A reference to the value being validated</param>
        /// <returns>The code to validate the reference of the given type</returns>
        public static string ValidateType(this IType type, IScopeProvider scope, string valueReference)
        {
            CompositeType model = type as CompositeType;
            SequenceType sequence = type as SequenceType;
            DictionaryType dictionary = type as DictionaryType;
            /*
            if (model != null && model.Properties.Any())
            {
                return CheckNotNull(valueReference, string.Format("{0}.Validate();", valueReference));
            }
            if (sequence != null)
            {
                var elementVar = scope.GetVariableName("element");
                var innerValidation = sequence.ElementType.ValidateType(scope, elementVar);
                if (!string.IsNullOrEmpty(innerValidation))
                {
                    var sequenceBuilder = string.Format("foreach ( var {0} in {1})\r\n{{\r\n", elementVar,
                        valueReference);
                    sequenceBuilder += string.Format("    {0}\r\n}}", innerValidation);
                    return CheckNotNull(valueReference, sequenceBuilder);
                }
            }
            else if (dictionary != null)
            {
                var valueVar = scope.GetVariableName("valueElement");
                var innerValidation = dictionary.ValueType.ValidateType(scope, valueVar);
                if (!string.IsNullOrEmpty(innerValidation))
                {
                    var sequenceBuilder = string.Format("if ( {0} != null)\r\n{{\r\n", valueReference);
                    sequenceBuilder += string.Format("    foreach ( var {0} in {1}.Values)\r\n    {{\r\n", valueVar,
                        valueReference);
                    sequenceBuilder += string.Format("        {0}\r\n    }}\r\n}}", innerValidation);
                    return CheckNotNull(valueReference, sequenceBuilder);
                }
            }
             */

            return null;
        }

        public static bool HasModelTypes(this Method method)
        {
            return method.Parameters.Any(p => p.Type is CompositeType) ||
                   method.Responses.Any(r => r.Value.Body is CompositeType) || method.ReturnType.Body is CompositeType ||
                   method.DefaultResponse.Body is CompositeType;
        }

        /////////////////////////////////////////////////////////////////////////////////////////
        //
        // ClientModel Extensions
        //
        /////////////////////////////////////////////////////////////////////////////////////////

        public static bool BelongsToGroup(this Method method, string groupName, string packageName)
        {
            return string.IsNullOrEmpty(groupName)
                    ? string.IsNullOrEmpty(method.Group) || string.Equals(method.Group, packageName, System.StringComparison.InvariantCultureIgnoreCase)
                    : string.Equals(method.Group, groupName);
        }

        public static bool HasReturnValue(this Method method)
        {
            return method.ReturnValue() != null && method.ReturnValue().Body != null;
        }

        public static Response ReturnValue(this Method method)
        {
            return method.ReturnType != null
                    ? method.ReturnType
                    : method.DefaultResponse;
        }
        
        public static bool IsPageable(this Method method)
        {
            return !string.IsNullOrEmpty(method.NextLink());
        }

        public static bool IsLongRunningOperation(this Method method)
        {
            try {
                return method.Extensions.ContainsKey(AzureExtensions.LongRunningExtension) && (bool)method.Extensions[AzureExtensions.LongRunningExtension];
            }
            catch(InvalidCastException e)
            {
                throw new Exception(string.Format(e.Message+" The value \'{0}\' for extension {1} for method {2} is invalid in Swagger. It should be boolean.",
                            method.Extensions[AzureExtensions.LongRunningExtension], AzureExtensions.LongRunningExtension, method.Group + "." + method.Name));
            }
        }

        public static string NextLink(this Method method)
        {
            var nextLink = "";

            // Note:
            // -- The CSharp generator applies a default link name if the extension is present but the link name is not.
            //    Yet, the MSDN for methods whose nextLink is missing are not paged methods. It appears the CSharp code is
            //    outdated vis a vis the specification.
            // TODO (gosdk): Ensure obtaining the nextLink is correct.
            if (method.Extensions.ContainsKey(AzureExtensions.PageableExtension))
            {
                var pageableExtension = method.Extensions[AzureExtensions.PageableExtension] as Newtonsoft.Json.Linq.JContainer;
                if (pageableExtension != null)
                {
                    var nextLinkName = (string)pageableExtension["nextLinkName"];
                    if (!string.IsNullOrEmpty(nextLinkName))
                    {
                        nextLink = GoCodeNamer.PascalCase(nextLinkName);
                    }
                }
            }

            return nextLink;
        }

    }
}
