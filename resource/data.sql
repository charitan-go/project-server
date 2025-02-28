-- Ensure the pgcrypto extension is enabled (for gen_random_uuid(), if used elsewhere)
--CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO projects (
    id, 
    readable_id,
    name,
    description,
    goal,
    start_date,
    end_date,
    category,
    country_code,
    status,
    owner_charity_readable_id
) VALUES
(
    '4f77e2b0-e686-459c-922a-aa333234c90b',
    '8922c6fd-ed02-4f2c-bd6b-2737eeed2120',
    'Middle East Crisis Relief',
    'A humanitarian relief initiative designed to provide critical support and aid to communities in crisis in the Middle East, organized by an Israeli non-profit organization.',
    50000.00,
    '2025-06-15 00:00:00',
    '2025-12-15 23:59:59',
    'HUMANTARIAN',
    'IL',
    'PENDING',
    'cf81bbc7-8990-4b9c-847d-ea0f3bc3e28a'
),
(
    '56acccc2-d6a0-4a38-9b3d-437f8d91c0c7',
    '6e48b458-0523-4f68-997d-11d2fdc7db2d',
    'Ukraine Medical Aid Initiative',
    'A comprehensive medical aid initiative aimed at providing emergency health services and relief to communities affected by the Russia-Ukraine conflict.',
    75000.00,
    '2025-07-01 00:00:00',
    '2025-12-31 23:59:59',
    'HEALTH',
    'UA',
    'PENDING',
    '5b08c663-a605-43e1-86dd-0941021f6383'
),
(
    'a11394b5-1bb0-4161-bb8a-99fbd51621e1',
    '1e6ffa4c-d523-4aa0-8ebe-74c95aeaef5b',
    'South Africa Food Program',
    'A sustainable food program focused on addressing hunger and improving nutrition for vulnerable communities in South Africa through community-based initiatives.',
    60000.00,
    '2025-08-01 00:00:00',
    '2025-10-31 23:59:59',
    'FOOD',
    'ZA',
    'PENDING',
    '921a99c0-47cd-4507-bc7c-bc284d15c7b9'
);

INSERT INTO projects (
    id,
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
    'e9e670b1-7e36-4831-82d2-6550be7fe29e',
    '3a64e46b-2027-433b-b06e-8a92828f6f3b',
    'Yagi Typhoon Support',
    'A humanitarian relief project providing emergency support to communities affected by typhoons in Vietnam, organized by an individual.',
    50000.00,
    '2025-07-01 00:00:00',
    '2025-12-31 23:59:59',
    'HUMANTARIAN',
    'VN',
    'PENDING',
    '47fe1f6d-05c4-4ff4-9e5d-91f47b585380'
),
(
    'ef9d5bad-1813-4fcc-ade8-2eaf29578325',
    '2c96beed-f78c-4c9d-a318-879969f8e78b',
    'Milton Hurricane Support',
    'A disaster recovery project supporting communities impacted by hurricanes in the USA, organized by an individual.',
    60000.00,
    '2025-08-01 00:00:00',
    '2026-01-31 23:59:59',
    'HUMANTARIAN',
    'US',
    'PENDING',
    'f61746df-3de1-47bd-88ee-3705df97dbac'
),
(
    '9fd8d2cc-97e4-4427-847f-e5bf06b28784',
    '5f86d71f-574c-445c-9886-e8ad44e00ec9',
    'Helping Ukrainian Refugee',
    'A humanitarian initiative aimed at providing critical aid and support to Ukrainian refugees in Germany, organized by a company.',
    70000.00,
    '2025-09-01 00:00:00',
    '2026-02-28 23:59:59',
    'HUMANTARIAN',
    'DE',
    'PENDING',
    'company'
    'c0132204-680d-450b-9fce-6f4f91f90dcf'
),
(
    '411ceab6-270b-4c03-8243-e28b5a40f1a5',
    '52ada42d-4a87-40eb-abe7-382bb3fd9a48',
    'Supporting SOS Children’s Village',
    'A project dedicated to supporting SOS Children’s Village in China by providing educational and housing support, organized by a company.',
    80000.00,
    '2025-10-01 00:00:00',
    '2026-03-31 23:59:59',
    'HUMANTARIAN',
    'CN',
    'PENDING',
    'c0132204-680d-450b-9fce-6f4f91f90dcf'
);
