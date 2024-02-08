// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationFromTemplateArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationFromTemplateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates an application registration and associated service principal from a gallery template.
 * 
 * &gt; The azuread.Application resource can also be used to instantiate a gallery application, however unlike the `azuread.Application` resource, this resource does not attempt to manage any properties of the resulting application.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.inputs.GetApplicationTemplateArgs;
 * import com.pulumi.azuread.ApplicationFromTemplate;
 * import com.pulumi.azuread.ApplicationFromTemplateArgs;
 * import com.pulumi.azuread.inputs.GetApplicationArgs;
 * import com.pulumi.azuread.inputs.GetServicePrincipalArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var exampleApplicationTemplate = AzureadFunctions.getApplicationTemplate(GetApplicationTemplateArgs.builder()
 *             .displayName(&#34;Marketo&#34;)
 *             .build());
 * 
 *         var exampleApplicationFromTemplate = new ApplicationFromTemplate(&#34;exampleApplicationFromTemplate&#34;, ApplicationFromTemplateArgs.builder()        
 *             .displayName(&#34;Example Application&#34;)
 *             .templateId(exampleApplicationTemplate.applyValue(getApplicationTemplateResult -&gt; getApplicationTemplateResult.templateId()))
 *             .build());
 * 
 *         final var exampleApplication = AzureadFunctions.getApplication(GetApplicationArgs.builder()
 *             .objectId(exampleApplicationFromTemplate.applicationObjectId())
 *             .build());
 * 
 *         final var exampleServicePrincipal = AzureadFunctions.getServicePrincipal(GetServicePrincipalArgs.builder()
 *             .objectId(exampleApplicationFromTemplate.servicePrincipalObjectId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Templated Applications can be imported using the template ID, the object ID of the application, and the object ID of the service principal, in the following format.
 * 
 * ```sh
 * $ pulumi import azuread:index/applicationFromTemplate:ApplicationFromTemplate example /applicationTemplates/00000000-0000-0000-0000-000000000000/instantiate/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
 * ```
 * 
 */
@ResourceType(type="azuread:index/applicationFromTemplate:ApplicationFromTemplate")
public class ApplicationFromTemplate extends com.pulumi.resources.CustomResource {
    /**
     * The resource ID for the application.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return The resource ID for the application.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * The object ID for the application.
     * 
     */
    @Export(name="applicationObjectId", refs={String.class}, tree="[0]")
    private Output<String> applicationObjectId;

    /**
     * @return The object ID for the application.
     * 
     */
    public Output<String> applicationObjectId() {
        return this.applicationObjectId;
    }
    /**
     * The display name for the application.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name for the application.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The resource ID for the service principal.
     * 
     */
    @Export(name="servicePrincipalId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalId;

    /**
     * @return The resource ID for the service principal.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }
    /**
     * The object ID for the service principal.
     * 
     */
    @Export(name="servicePrincipalObjectId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalObjectId;

    /**
     * @return The object ID for the service principal.
     * 
     */
    public Output<String> servicePrincipalObjectId() {
        return this.servicePrincipalObjectId;
    }
    /**
     * Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="templateId", refs={String.class}, tree="[0]")
    private Output<String> templateId;

    /**
     * @return Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> templateId() {
        return this.templateId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationFromTemplate(String name) {
        this(name, ApplicationFromTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationFromTemplate(String name, ApplicationFromTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationFromTemplate(String name, ApplicationFromTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationFromTemplate:ApplicationFromTemplate", name, args == null ? ApplicationFromTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApplicationFromTemplate(String name, Output<String> id, @Nullable ApplicationFromTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationFromTemplate:ApplicationFromTemplate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ApplicationFromTemplate get(String name, Output<String> id, @Nullable ApplicationFromTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationFromTemplate(name, id, state, options);
    }
}
