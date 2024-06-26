// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrivilegedAccessGroupEligibilityScheduleState extends com.pulumi.resources.ResourceArgs {

    public static final PrivilegedAccessGroupEligibilityScheduleState Empty = new PrivilegedAccessGroupEligibilityScheduleState();

    /**
     * The type of assignment to the group. Can be either `member` or `owner`.
     * 
     */
    @Import(name="assignmentType")
    private @Nullable Output<String> assignmentType;

    /**
     * @return The type of assignment to the group. Can be either `member` or `owner`.
     * 
     */
    public Optional<Output<String>> assignmentType() {
        return Optional.ofNullable(this.assignmentType);
    }

    /**
     * The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
     * 
     */
    @Import(name="duration")
    private @Nullable Output<String> duration;

    /**
     * @return The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
     * 
     */
    public Optional<Output<String>> duration() {
        return Optional.ofNullable(this.duration);
    }

    /**
     * The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
     * 
     */
    @Import(name="expirationDate")
    private @Nullable Output<String> expirationDate;

    /**
     * @return The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
     * 
     */
    public Optional<Output<String>> expirationDate() {
        return Optional.ofNullable(this.expirationDate);
    }

    /**
     * The Object ID of the Azure AD group to which the principal will be assigned.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    /**
     * @return The Object ID of the Azure AD group to which the principal will be assigned.
     * 
     */
    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * The justification for this assignment. May be required by the role policy.
     * 
     */
    @Import(name="justification")
    private @Nullable Output<String> justification;

    /**
     * @return The justification for this assignment. May be required by the role policy.
     * 
     */
    public Optional<Output<String>> justification() {
        return Optional.ofNullable(this.justification);
    }

    /**
     * Is this assigment permanently valid.
     * 
     * At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
     * 
     */
    @Import(name="permanentAssignment")
    private @Nullable Output<Boolean> permanentAssignment;

    /**
     * @return Is this assigment permanently valid.
     * 
     * At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
     * 
     */
    public Optional<Output<Boolean>> permanentAssignment() {
        return Optional.ofNullable(this.permanentAssignment);
    }

    /**
     * The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
     * 
     */
    @Import(name="principalId")
    private @Nullable Output<String> principalId;

    /**
     * @return The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
     * 
     */
    public Optional<Output<String>> principalId() {
        return Optional.ofNullable(this.principalId);
    }

    /**
     * The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
     * 
     */
    @Import(name="startDate")
    private @Nullable Output<String> startDate;

    /**
     * @return The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
     * 
     */
    public Optional<Output<String>> startDate() {
        return Optional.ofNullable(this.startDate);
    }

    /**
     * (String) The provisioning status of this request.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return (String) The provisioning status of this request.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The ticket number in the ticket system approving this assignment. May be required by the role policy.
     * 
     */
    @Import(name="ticketNumber")
    private @Nullable Output<String> ticketNumber;

    /**
     * @return The ticket number in the ticket system approving this assignment. May be required by the role policy.
     * 
     */
    public Optional<Output<String>> ticketNumber() {
        return Optional.ofNullable(this.ticketNumber);
    }

    /**
     * The ticket system containing the ticket number approving this assignment. May be required by the role policy.
     * 
     */
    @Import(name="ticketSystem")
    private @Nullable Output<String> ticketSystem;

    /**
     * @return The ticket system containing the ticket number approving this assignment. May be required by the role policy.
     * 
     */
    public Optional<Output<String>> ticketSystem() {
        return Optional.ofNullable(this.ticketSystem);
    }

    private PrivilegedAccessGroupEligibilityScheduleState() {}

    private PrivilegedAccessGroupEligibilityScheduleState(PrivilegedAccessGroupEligibilityScheduleState $) {
        this.assignmentType = $.assignmentType;
        this.duration = $.duration;
        this.expirationDate = $.expirationDate;
        this.groupId = $.groupId;
        this.justification = $.justification;
        this.permanentAssignment = $.permanentAssignment;
        this.principalId = $.principalId;
        this.startDate = $.startDate;
        this.status = $.status;
        this.ticketNumber = $.ticketNumber;
        this.ticketSystem = $.ticketSystem;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivilegedAccessGroupEligibilityScheduleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivilegedAccessGroupEligibilityScheduleState $;

        public Builder() {
            $ = new PrivilegedAccessGroupEligibilityScheduleState();
        }

        public Builder(PrivilegedAccessGroupEligibilityScheduleState defaults) {
            $ = new PrivilegedAccessGroupEligibilityScheduleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param assignmentType The type of assignment to the group. Can be either `member` or `owner`.
         * 
         * @return builder
         * 
         */
        public Builder assignmentType(@Nullable Output<String> assignmentType) {
            $.assignmentType = assignmentType;
            return this;
        }

        /**
         * @param assignmentType The type of assignment to the group. Can be either `member` or `owner`.
         * 
         * @return builder
         * 
         */
        public Builder assignmentType(String assignmentType) {
            return assignmentType(Output.of(assignmentType));
        }

        /**
         * @param duration The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
         * 
         * @return builder
         * 
         */
        public Builder duration(@Nullable Output<String> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
         * 
         * @return builder
         * 
         */
        public Builder duration(String duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param expirationDate The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(@Nullable Output<String> expirationDate) {
            $.expirationDate = expirationDate;
            return this;
        }

        /**
         * @param expirationDate The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(String expirationDate) {
            return expirationDate(Output.of(expirationDate));
        }

        /**
         * @param groupId The Object ID of the Azure AD group to which the principal will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The Object ID of the Azure AD group to which the principal will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param justification The justification for this assignment. May be required by the role policy.
         * 
         * @return builder
         * 
         */
        public Builder justification(@Nullable Output<String> justification) {
            $.justification = justification;
            return this;
        }

        /**
         * @param justification The justification for this assignment. May be required by the role policy.
         * 
         * @return builder
         * 
         */
        public Builder justification(String justification) {
            return justification(Output.of(justification));
        }

        /**
         * @param permanentAssignment Is this assigment permanently valid.
         * 
         * At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
         * 
         * @return builder
         * 
         */
        public Builder permanentAssignment(@Nullable Output<Boolean> permanentAssignment) {
            $.permanentAssignment = permanentAssignment;
            return this;
        }

        /**
         * @param permanentAssignment Is this assigment permanently valid.
         * 
         * At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
         * 
         * @return builder
         * 
         */
        public Builder permanentAssignment(Boolean permanentAssignment) {
            return permanentAssignment(Output.of(permanentAssignment));
        }

        /**
         * @param principalId The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
         * 
         * @return builder
         * 
         */
        public Builder principalId(@Nullable Output<String> principalId) {
            $.principalId = principalId;
            return this;
        }

        /**
         * @param principalId The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
         * 
         * @return builder
         * 
         */
        public Builder principalId(String principalId) {
            return principalId(Output.of(principalId));
        }

        /**
         * @param startDate The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
         * 
         * @return builder
         * 
         */
        public Builder startDate(@Nullable Output<String> startDate) {
            $.startDate = startDate;
            return this;
        }

        /**
         * @param startDate The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
         * 
         * @return builder
         * 
         */
        public Builder startDate(String startDate) {
            return startDate(Output.of(startDate));
        }

        /**
         * @param status (String) The provisioning status of this request.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status (String) The provisioning status of this request.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param ticketNumber The ticket number in the ticket system approving this assignment. May be required by the role policy.
         * 
         * @return builder
         * 
         */
        public Builder ticketNumber(@Nullable Output<String> ticketNumber) {
            $.ticketNumber = ticketNumber;
            return this;
        }

        /**
         * @param ticketNumber The ticket number in the ticket system approving this assignment. May be required by the role policy.
         * 
         * @return builder
         * 
         */
        public Builder ticketNumber(String ticketNumber) {
            return ticketNumber(Output.of(ticketNumber));
        }

        /**
         * @param ticketSystem The ticket system containing the ticket number approving this assignment. May be required by the role policy.
         * 
         * @return builder
         * 
         */
        public Builder ticketSystem(@Nullable Output<String> ticketSystem) {
            $.ticketSystem = ticketSystem;
            return this;
        }

        /**
         * @param ticketSystem The ticket system containing the ticket number approving this assignment. May be required by the role policy.
         * 
         * @return builder
         * 
         */
        public Builder ticketSystem(String ticketSystem) {
            return ticketSystem(Output.of(ticketSystem));
        }

        public PrivilegedAccessGroupEligibilityScheduleState build() {
            return $;
        }
    }

}
