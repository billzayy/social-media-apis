package repositories

var getUserByIdQuery = `SELECT u."ID", u."UserName", u."FullName", u."Email", u."UserSince",
		u."LastLogin",u."ProfilePicture", u."Active",
		up."Description", up."Location", up."Website", up."BirthDate", 
		us."Theme", us."Language", us."Country" FROM public."Users" u 
	LEFT JOIN public."UserProfile" up ON up."UserId" = u."ID" 
	LEFT JOIN public."UserSettings" us ON us."UserId" = u."ID"
	WHERE u."ID" = $1 limit 1000;`

var updateUserTable = `
-- Update Users table
UPDATE public."Users"
SET 
    "UserName" = COALESCE($1, "UserName"),
    "Email" = COALESCE($2, "Email"),
    "FullName" = COALESCE($3, "FullName"),
    "ProfilePicture" = COALESCE($4, "ProfilePicture")
WHERE "ID" = $5;
`

var updateUserProfile = `
-- Update UserProfile table
UPDATE public."UserProfile"
SET
    "Description" = COALESCE($1, "Description"),
    "Location" = COALESCE($2, "Location"),
    "Website" = COALESCE($3, "Website"),
    "BirthDate" = COALESCE($4, "BirthDate")
WHERE "UserId" = $5;
`

var updateUserSetting = `
-- Update UserSettings table
UPDATE public."UserSettings"
SET
    "Theme" = COALESCE($1, "Theme"),
    "Language" = COALESCE($2, "Language"),
    "Country" = COALESCE($3, "Country")
WHERE "UserId" = $4;
`

var getPasswordQuery = `SELECT "Password" FROM public."Users" WHERE "ID" = $1`
var updatePassword = `UPDATE public."Users" SET "Password" = $1 WHERE "ID" = $2`
