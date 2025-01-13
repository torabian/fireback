package workspaces

// VisibilityLevel represents the various visibility levels for an entity.
// These levels determine who can access the data and under what conditions.
type VisibilityLevel uint8

const (
	// VISIBILITY_DEFAULT: Default visibility (0) where data is visible to the entire workspace
	// and the person who created it.
	// Use case: General workspace data like shared resources or common tasks visible to all members.
	VISIBILITY_DEFAULT VisibilityLevel = 0

	// VISIBILITY_PUBLIC: Logged-in users can see the record.
	// Use case: Content meant for authenticated users, such as user dashboards or community posts.
	VISIBILITY_PUBLIC VisibilityLevel = 1

	// VISIBILITY_OWNER: Only the owner of the record can see it.
	// Use case: Personal data or user-specific preferences that should remain private to the owner.
	VISIBILITY_OWNER VisibilityLevel = 2

	// VISIBILITY_PRIVATE: Only workspace members can see the record.
	// Use case: Internal workspace content like team-specific files or tasks.
	VISIBILITY_PRIVATE VisibilityLevel = 3

	// VISIBILITY_ANONYMOUS: Anyone can see the record, no authentication required.
	// Use case: Publicly available content like blog posts, marketing pages, or open-access data.
	VISIBILITY_ANONYMOUS VisibilityLevel = 4

	// VISIBILITY_ADMIN_ONLY: Only administrators of the workspace can see the record.
	// Use case: Highly sensitive data like audit logs, billing information, or administrative settings.
	VISIBILITY_ADMIN_ONLY VisibilityLevel = 5

	// VISIBILITY_GROUP: Only specific user groups can see the record.
	// Use case: Role-based access for features like project-specific files visible to certain teams.
	VISIBILITY_GROUP VisibilityLevel = 6

	// VISIBILITY_TIME_LIMITED: Visible only for a specific time period.
	// Use case: Temporary promotions, expiring content, or time-sensitive offers.
	VISIBILITY_TIME_LIMITED VisibilityLevel = 7

	// VISIBILITY_GEO_RESTRICTED: Accessible only from certain geographic locations.
	// Use case: Geo-fenced content like region-specific promotions or country-specific compliance data.
	VISIBILITY_GEO_RESTRICTED VisibilityLevel = 8

	// VISIBILITY_READ_ONLY: Visible to everyone but cannot be modified.
	// Use case: Historical or archived data that should be accessible for reference but not editable.
	VISIBILITY_READ_ONLY VisibilityLevel = 9
)
