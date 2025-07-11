package repositories

var getNotifyDataByIdQuery = `
SELECT "Id", "SenderId", "ReceiverId", "Messages", "Type", "Url", "IsRead", "Date"
FROM public."UserNotifications"
WHERE "ReceiverId" = $1`

var getUserSettingQuery = `
SELECT "UserId", "EmailNotifications", "PushNotifications" 
FROM public."UserSettings"
WHERE "UserId" = $1`

var getUnreadNotify = `
SELECT COUNT("Id") FROM public."UserNotifications" WHERE "ReceiverId" = $1 AND "IsRead" = false`

var saveNotifyQuery = `
INSERT INTO public."UserNotifications" ("SenderId","ReceiverId","Messages", "Type", "Url")
VALUES ($1,$2,$3,$4,$5)`

var updateRead = `UPDATE public."UserNotifications" SET "IsRead" = true WHERE "Id" = $1`
