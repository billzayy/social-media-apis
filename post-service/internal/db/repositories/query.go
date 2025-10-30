package repositories

var GetPostQuery = `
	SELECT p."ID" AS "PostId",
		json_build_object(
			'ID', u."ID",
			'FullName', u."FullName",
			'ProfilePicture', u."ProfilePicture"
		) AS "userId", p."Content", p."CreatedAt",
		COALESCE(
			json_agg(
				json_build_object(
					'PostId', pm."PostId",
					'Type', pm."Type",
					'Url', pm."Url"
				)
			) FILTER (WHERE pm."PostId" IS NOT NULL),
			'[]'
		) AS "media",
		(SELECT COUNT(*) FROM public."PostLikes" pl WHERE pl."PostId" = p."ID") AS "likes",
		(SELECT COUNT(*) FROM public."PostComments" pc WHERE pc."PostId" = p."ID") AS "comments"
	FROM public."Posts" p
	LEFT JOIN public."Users" u ON u."ID" = p."UserId"
	LEFT JOIN public."PostMedia" pm ON pm."PostId" = p."ID"
	GROUP BY p."ID", u."ID", u."FullName", u."ProfilePicture"
	ORDER BY p."CreatedAt" DESC
LIMIT 1000;`
