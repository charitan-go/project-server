INSERT INTO projects (
    readable_id,
    name,
    description,
    goal,
    start_date,
    end_date,
    category,
    countryCode,
    status,
    owner_charity_readable_id
) VALUES
(
    '8922c6fd-ed02-4f2c-bd6b-2737eeed2120',
    'Middle East Crisis Relief',
    'A humanitarian relief initiative designed to provide critical support and aid to communities in crisis in the Middle East, organized by an Israeli non-profit organization.',
    50000.00,
    '2025-06-15 00:00:00',
    '2025-12-15 23:59:59',
    'HUMANTARIAN',
    'IL',
    'PENDING',
    'owner-charity-A'
),
(
    '6e48b458-0523-4f68-997d-11d2fdc7db2d',
    'Ukraine Medical Aid Initiative',
    'A comprehensive medical aid initiative aimed at providing emergency health services and relief to communities affected by the Russia-Ukraine conflict.',
    75000.00,
    '2025-07-01 00:00:00',
    '2025-12-31 23:59:59',
    'HEALTH',
    'UA',
    'PENDING',
    'owner-charity-B'
),
(
    '1e6ffa4c-d523-4aa0-8ebe-74c95aeaef5b',
    'South Africa Food Program',
    'A sustainable food program focused on addressing hunger and improving nutrition for vulnerable communities in South Africa through community-based initiatives.',
    60000.00,
    '2025-08-01 00:00:00',
    '2025-10-31 23:59:59',
    'FOOD',
    'ZA',
    'PENDING',
    'owner-charity-C'
);
