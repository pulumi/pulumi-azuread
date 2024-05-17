// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.SynchronizationJobProvisionOnDemandArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.SynchronizationJobProvisionOnDemandState;
import com.pulumi.azuread.outputs.SynchronizationJobProvisionOnDemandParameter;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages synchronization job on demand provisioning associated with a service principal (enterprise application) within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `Synchronization.ReadWrite.All`
 * 
 * ## Example Usage
 * 
 * *Basic example*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.Group;
 * import com.pulumi.azuread.GroupArgs;
 * import com.pulumi.azuread.inputs.GetApplicationTemplateArgs;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.inputs.ApplicationFeatureTagArgs;
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import com.pulumi.azuread.SynchronizationSecret;
 * import com.pulumi.azuread.SynchronizationSecretArgs;
 * import com.pulumi.azuread.inputs.SynchronizationSecretCredentialArgs;
 * import com.pulumi.azuread.SynchronizationJob;
 * import com.pulumi.azuread.SynchronizationJobArgs;
 * import com.pulumi.azuread.SynchronizationJobProvisionOnDemand;
 * import com.pulumi.azuread.SynchronizationJobProvisionOnDemandArgs;
 * import com.pulumi.azuread.inputs.SynchronizationJobProvisionOnDemandParameterArgs;
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
 *         final var current = AzureadFunctions.getClientConfig();
 * 
 *         var exampleGroup = new Group("exampleGroup", GroupArgs.builder()        
 *             .displayName("example")
 *             .owners(current.applyValue(getClientConfigResult -> getClientConfigResult.objectId()))
 *             .securityEnabled(true)
 *             .build());
 * 
 *         final var example = AzureadFunctions.getApplicationTemplate(GetApplicationTemplateArgs.builder()
 *             .displayName("Azure Databricks SCIM Provisioning Connector")
 *             .build());
 * 
 *         var exampleApplication = new Application("exampleApplication", ApplicationArgs.builder()        
 *             .displayName("example")
 *             .templateId(example.applyValue(getApplicationTemplateResult -> getApplicationTemplateResult.templateId()))
 *             .featureTags(ApplicationFeatureTagArgs.builder()
 *                 .enterprise(true)
 *                 .gallery(true)
 *                 .build())
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal("exampleServicePrincipal", ServicePrincipalArgs.builder()        
 *             .clientId(exampleApplication.clientId())
 *             .useExisting(true)
 *             .build());
 * 
 *         var exampleSynchronizationSecret = new SynchronizationSecret("exampleSynchronizationSecret", SynchronizationSecretArgs.builder()        
 *             .servicePrincipalId(exampleServicePrincipal.id())
 *             .credentials(            
 *                 SynchronizationSecretCredentialArgs.builder()
 *                     .key("BaseAddress")
 *                     .value("https://adb-example.azuredatabricks.net/api/2.0/preview/scim")
 *                     .build(),
 *                 SynchronizationSecretCredentialArgs.builder()
 *                     .key("SecretToken")
 *                     .value("some-token")
 *                     .build())
 *             .build());
 * 
 *         var exampleSynchronizationJob = new SynchronizationJob("exampleSynchronizationJob", SynchronizationJobArgs.builder()        
 *             .servicePrincipalId(exampleServicePrincipal.id())
 *             .templateId("dataBricks")
 *             .enabled(true)
 *             .build());
 * 
 *         var exampleSynchronizationJobProvisionOnDemand = new SynchronizationJobProvisionOnDemand("exampleSynchronizationJobProvisionOnDemand", SynchronizationJobProvisionOnDemandArgs.builder()        
 *             .servicePrincipalId(exampleServicePrincipal.id())
 *             .synchronizationJobId(exampleSynchronizationJob.id())
 *             .parameters(SynchronizationJobProvisionOnDemandParameterArgs.builder()
 *                 .ruleId("")
 *                 .subjects(SynchronizationJobProvisionOnDemandParameterSubjectArgs.builder()
 *                     .objectId(exampleGroup.objectId())
 *                     .objectTypeName("Group")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource does not support importing.
 * 
 */
@ResourceType(type="azuread:index/synchronizationJobProvisionOnDemand:SynchronizationJobProvisionOnDemand")
public class SynchronizationJobProvisionOnDemand extends com.pulumi.resources.CustomResource {
    /**
     * One or more `parameter` blocks as documented below.
     * 
     */
    @Export(name="parameters", refs={List.class,SynchronizationJobProvisionOnDemandParameter.class}, tree="[0,1]")
    private Output<List<SynchronizationJobProvisionOnDemandParameter>> parameters;

    /**
     * @return One or more `parameter` blocks as documented below.
     * 
     */
    public Output<List<SynchronizationJobProvisionOnDemandParameter>> parameters() {
        return this.parameters;
    }
    /**
     * The object ID of the service principal for the synchronization job.
     * 
     */
    @Export(name="servicePrincipalId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalId;

    /**
     * @return The object ID of the service principal for the synchronization job.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }
    /**
     * Identifier of the synchronization template this job is based on.
     * 
     */
    @Export(name="synchronizationJobId", refs={String.class}, tree="[0]")
    private Output<String> synchronizationJobId;

    /**
     * @return Identifier of the synchronization template this job is based on.
     * 
     */
    public Output<String> synchronizationJobId() {
        return this.synchronizationJobId;
    }
    @Export(name="triggers", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> triggers;

    public Output<Optional<Map<String,String>>> triggers() {
        return Codegen.optional(this.triggers);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SynchronizationJobProvisionOnDemand(String name) {
        this(name, SynchronizationJobProvisionOnDemandArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SynchronizationJobProvisionOnDemand(String name, SynchronizationJobProvisionOnDemandArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SynchronizationJobProvisionOnDemand(String name, SynchronizationJobProvisionOnDemandArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/synchronizationJobProvisionOnDemand:SynchronizationJobProvisionOnDemand", name, args == null ? SynchronizationJobProvisionOnDemandArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SynchronizationJobProvisionOnDemand(String name, Output<String> id, @Nullable SynchronizationJobProvisionOnDemandState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/synchronizationJobProvisionOnDemand:SynchronizationJobProvisionOnDemand", name, state, makeResourceOptions(options, id));
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
    public static SynchronizationJobProvisionOnDemand get(String name, Output<String> id, @Nullable SynchronizationJobProvisionOnDemandState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SynchronizationJobProvisionOnDemand(name, id, state, options);
    }
}