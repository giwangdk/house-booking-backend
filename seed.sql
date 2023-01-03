INSERT INTO cities VALUES 
(1, 'Aceh'),
(2, 'Medan'),
(3, 'Palembang'),
(4, 'Bandung'),
(5, 'Jakarta Timur'),
(6, 'Jakarta Selatan'),
(7, 'Jakarta Barat'),
(8, 'Jakarta Utara'),
(9, 'Jakarta Pusat'),
(10, 'Surabaya'),
(11, 'Bali'),
(12, 'Yogyakarta'),
(13, 'Lombok'),
(14, 'Bogor'),
(15, 'Semarang'),
(16, 'Banten'),
(17, 'Lampung'),
(18, 'Bengkulu'),
(19, 'Jambi'),
(20, 'Riau'),
(21, 'Bengkulu'),
(22, 'Jambi'),
(23, 'Lampung'),
(24, 'Banten'),
(25, 'Bandung'),
(26, 'Sukabumi'),
(27, 'Cirebon'),
(28, 'Tasikmalaya'),
(29, 'Garut'),
(30, 'Ciamis'),
(31, 'Kuningan'),
(32, 'Cianjur'),
(33, 'Majalengka'),
(34, 'Sumedang'),
(35, 'Indramayu'),
(36, 'Subang'),
(37, 'Purwakarta'),
(38, 'Karawang'),
(39, 'Bekasi');

INSERT INTO reservation_status VALUES 
(1, 'PENDING'),
(2, 'SETTLEMENT'),
(3, 'FAILURE'),
(4, 'WAITING CONFIRMATION');

INSERT INTO pickup_statuses VALUES 
(1, 'Pending Admin'),
(2, 'Awaiting Check-in Date'),
(3, 'On the Way Pickup'),
(4, 'On the Way Reservation'),
(5, 'Completed');

INSERT INTO users (email, fullname, password, address, role, city_id) VALUES 
('admin@gmail.com', 'admin giwang', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'Perum Baros', 'admin', '1'),
('host@gmail.com', 'host giwang', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'jl. Masjid Darussalam No. 11', 'host', '2'),
('user@gmail.com', 'user giwang', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'jl. Masjid Yerussalem No. 11', 'user', '4'),
('giwangdk@gmail.com', 'user 1', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'jl. Soeku Yerussalem No. 21', 'user', '3'),
('gidwikintan@gmail.com', 'user 2', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'jl. Penjaittan Yerussalem No. 01', 'user', '6'),
('host2@gmail.com', 'host2 giwang', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'jl. Masjid Darussalam No. 11', 'host', '10'),
('tifany@gmail.com', 'Tifany Angelia', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'Jl. Malang Buana', 'user', '21'),
('angelia@gmail.com', 'Tifany Angelia 2', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'Perum Bumi perumahan', 'user', '11'),
('sekar@gmail.com', 'Sekar Nityasa', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'Sekuran Jalan', 'user', '21'),
('cikal@gmail.com', 'Cikal Tania', '$2a$04$mBFHO40jsIJmrLFHRGV4s.e62YbJ58MgU2WQnoUGExFxfb2inxe2a', 'Jl. Leuwiliang', 'host', '7');



INSERT INTO wallets(balance, user_id) VALUES
( 1000000, 1),( 1000000, 2),( 1000000, 3),( 1000000, 4),( 1000000, 5), ( 10000000, 6), ( 1000000, 7), ( 10000000, 8), ( 1000000, 9), ( 10000000, 10);

INSERT INTO games (user_id, chance, total_games_played) VALUES
( 1, 10, 0),
( 2, 10, 0),
( 3, 10, 0),
( 4, 10, 0),
( 5, 10, 0),
( 6, 10, 0),
( 7, 10, 0),
( 8, 10, 0),
( 9, 10, 0),
( 10, 10, 0);


INSERT INTO houses (name, city_id, user_id, price, description, location) VALUES 
(
 'Aura House 2bds Eco Bamboo House, Pool, River View', 
 11,
 2,
3600000, 
'Aura house is a beautiful & unique eco bamboo house built on the west bank of the River Ayung facing east to catch sunrise. Aura House is situated 25min away from Ubud, and 35min away from Canggu.',
'Canggu'
),
(
 'Lespoir II - ocean view+huge pool+ butler', 
 11,
 2,
5600000, 
'Villa Lespoir II - Open on 20th Feb. 2020, our 2nd villa which is just about 40m away from Villa Lespoir I. Again luxury PANORAMIC view villa in tranquil northern Bali. It is at where the green vineyards and rice fields meet the ocean.',
'Seririt'
),
(
 'Villa Saluka - Romantic & Serene with Sunset View', 
 12,
 6,
3600000, 
'Dear future guests, thank you so much for coming across our page. A little bit about Villa Saluka: the name is from my lovely little family. SAtria (my husband), LUise (me), KAi (my son). I initially built this house for us, but for a few reasons, we couldn’t live there. ',
 'Bantul'
 ),
(
 'Ginza Bungalow Yogyakarta', 
 12,
 6,
3600000, 
'Located in the heart of Jogjakarta
Sleeps 4 adults 
King size bed and double mattress in loft 
Lap swim pool
Covered large entertaining deck areas',
 'Banguntapan'
 ),
(
 'Luxury 1 bedroom at parahyangan residence', 
4,
 6,
3600000, 
'Kick back and relax in this calm, stylish space.',
 'Cidadap'
 ),
(
 'Brand New Bright Studio @ Landmark City Centre', 
 25,
 2,
2600000, 
'Our bright and lively brand new studio apartment at lv 12 in Landmark Residence is located strategically in the heart of Bandung. You will be close to everything when you stay at this centrally-located place. The complex is one of Bandungs most verdant and prestigious residence. ',
 'Canggu'
 ),
(
 'Oribu Kidul, Direct Sunrise view of Mount Gede', 
 26,
 2,
4000000, 
'Relax, unwind, or even working. This tiny villa is a place for all. 
 Nestled in the feet of Mount Salak overviewing Mount Pangrango and our own garden. 
 Your solution for short getaway from citys hustle, 70 minutes away from Jakarta.',
 'Cigombong'
 ),
 (
 'Villa Pusat kota Batu 2br,3menit dari Alun2 Batu', 
 9,
 6,
100000, 
'Gaya industrial minimalis',
 'Bitu'
 ),
 (
 'RiceField Villa Bukit Lawang', 
 2,
 6,
150000, 
'It is a house of stone, bamboo and wood. Built 100% by hand with local materials in 2015. It is in the middle of a rice plantation in the rural area of Bukit Lawang and just 10 minutes walk from the bus station. Perfect for relaxing as a couple or with friends. It has mountain, sunset and sunrise views. A unique and authentic house in Bukit Lawang',
 'Bohorok'
 ),
(
 'New Home with own beach frontage', 
 26,
 6,
150000, 
'It has four Bedrooms and 2.5 Bathrooms and two large living areas opening onto huge decks to allow for relaxed outdoor beach-side living. Furnishing and decor is  modern and minimalist. Fully air-conditioned, fans, WiFi, TV and sound system.',
  'Pelabuhan Ratu'
),
(
 'Quiet Villa and Pool with Mt. and Lake Batur View', 
 11,
 6,
6500000,
'This villa is located in Batur Kintamani Bali. Located right at the foot of Mount Batur so you can enjoy the beauty of Mount Batur and Lake Batur. The villa environment is very comfortable, the atmosphere is very calm, and the air is still fresh and far from the crowds of the city. You can enjoy the beautiful views of the mountains and lakes while swimming.',
  'Bangli'
),
(
 '180° VIEW, PRIVATE POOL VILLA..', 
 5,
 2,
3500000,
'The villa sits in an elevated postion in one of the last remaining untouched parts of the island.',
  'Buyu'
  ),
  (
 'Villa Asi', 
 11,
 10,
7000000,
'This spectacular home on Koh Samui is situated in a quiet enclave along the east coast near Bophut and Chaweng. Perched on elevated land about five kilometers from the beach, the villa gazes out on panoramic views of Chaweng Bay. Its superb architectural design features tiered terraces cascading down the lush hillside, including a split-level infinity swimming pool fused by a central waterfall',
  'Phuket'
  ),
   (
 'Samujana Seventeen', 
 21,
 10,
4000000,
'Blurring indoor and outdoor spaces in the most spectacular fashion, this private luxury villa is the perfect atmosphere for enjoying Thailand’s beautiful ocean environment. Located on the hillside of the Samujana Estate, Villa Seventeen boasts some of the areas most spectacular ocean views.',
  'Phuket'
  ),
(
 'Bali Bamboo House | Rescape Ubud - Resound Villa', 
 11,
 10,
3000000,
'Rescape (retreat/ escape) Ubud is an uniquely designed villa built by bamboo, allowing guests to unwind and escape from all the daily hassle. This stay is perfect for couples, young families, artists, musicians and everyone who loves to fully embrace the nature. ',
  'Ubud'
  ),
(
 'Noku Beach House', 
 4,
 10,
3000000,
'A 30-foot infinity pool reaches toward the sea at this wood-clad villa on Seminyak Beach. A walled garden and attentive staff look onto the scene, with interiors by Alex Zabotto-Bentley adding potted plants, maritime sculptures, and designer furnishings. ',
  'Badunk'
  ),
(
 'Breakneck Gorge Oikos - Multi-Award Winning Luxury Retreat!', 
 22,
 10,
6000000,
'Set on a startlingly beautiful elevation on the sprawling 18-hectare property of Breakneck Gorge, and just minutes from the main street of Daylesford, Oikos was designed as an indulgent retreat from the city. ',
  'Brisbanr'
  ),
(
 'Regal Residence!', 
 9,
 2,
8000000,
'he beach, the bay, and Chania Town are only overshadowed by the ocean sunset in the panoramic vista from this amazing hilltop villa. Soak it all in from the swimming pool, while enjoying alfresco meals, or relaxing on a plush day bed. ',
  'Yuhi'
  ),
(
 'the naked house', 
 3,
 2,
3000000,
'Quarantining, social distancing and remote working is what this house is perfect for. Please ask us about special covid rates. It is an architectural villa on the south side of Koh Samui, private and in a natural environment, it has sweeping ocean views and has a great salt water lap pool. Half way up a hill, it gets natural breezes, whithout mozzies even at dusk. ',
  'Yuhi'
  ),
  (
 'Penthouse 180m from beachfront with pool and sauna', 
 17,
 2,
9000000,
'Slowly exhale at the end of a day and watch the sun set into the sea. Arise in the morning with 360 degrees views of the mountains and coastline. Walk to Camps Bay beach with all the top restaurants and outdoor bars. ',
  'Jiula'
  );






INSERT INTO house_photos(house_id, photo) VALUES
(1, 'https://a0.muscache.com/im/pictures/e25a9b25-fa98-4160-bfd1-039287bf38b6.jpg?im_w=1200'),
(1,'https://a0.muscache.com/im/pictures/miso/Hosting-34113796/original/f4f7b242-db33-46fc-9080-c3d6a6fd55ec.jpeg?im_w=1440'),
(1,'https://a0.muscache.com/im/pictures/miso/Hosting-34113796/original/372e7d6f-7fb9-4668-92f0-25bb9a346814.jpeg?im_w=1440'),
(1, 'https://a0.muscache.com/im/pictures/miso/Hosting-34113796/original/4756e699-f474-4ca7-8b77-06b12715a6cb.jpeg?im_w=1440'),
(1,'https://a0.muscache.com/im/pictures/miso/Hosting-34113796/original/fca892a4-3481-4ad1-9f92-404feaa42e9f.jpeg?im_w=1440'),
(1,'https://a0.muscache.com/im/pictures/miso/Hosting-34113796/original/36d8007a-0de5-439d-9fec-1c2d7b53a147.jpeg?im_w=1440'),
(1, 'https://a0.muscache.com/im/pictures/miso/Hosting-34113796/original/f95b0a2e-0272-469e-a56c-433b9cc4ffdb.jpeg?im_w=1440'),
(1, 'https://a0.muscache.com/im/pictures/miso/Hosting-34113796/original/97108083-280f-4e0e-9f1b-7a4b0dd34c44.jpeg?im_w=1440'),
(1, 'https://a0.muscache.com/im/pictures/miso/Hosting-34113796/original/8a9f650e-be60-47a9-833f-2ed0950cb9f9.jpeg?im_w=1440'),
(1, 'https://a0.muscache.com/im/pictures/miso/Hosting-34113796/original/6308de24-4c50-4bb6-8a66-ae42b45a2ac2.jpeg?im_w=1440'),
(2, 'https://a0.muscache.com/im/pictures/0c654870-50d9-4d1c-ba1d-f50b6e6cbdd5.jpg?im_w=1200'),
(2, 'https://a0.muscache.com/im/pictures/c1b03e42-9d2b-4c9b-8411-8ed4feee4907.jpg?im_w=1200'),
(2, 'https://a0.muscache.com/im/pictures/9a3a059a-2061-44b0-98dd-43a7c11812cc.jpg?im_w=1200'),
(2, 'https://a0.muscache.com/im/pictures/0d88a4e0-017b-4806-bc0e-2a31f68e7533.jpg?im_w=1200'),
(2, 'https://a0.muscache.com/im/pictures/15637107-d330-4c2d-b120-5faa69cc3e5e.jpg?im_w=1200'),
(2, 'https://a0.muscache.com/im/pictures/15cffe98-78bc-4498-98bd-af94522038e4.jpg?im_w=1200'),
(2, 'https://a0.muscache.com/im/pictures/1f12bc12-f137-4da4-8d35-54317d1bc15d.jpg?im_w=1200'),
(2, 'https://a0.muscache.com/im/pictures/6a94371e-79c0-4234-a372-44225dca764f.jpg?im_w=1200'),
(3, 'https://a0.muscache.com/im/pictures/112103a3-f197-45a8-aa33-1d8774da2e4c.jpg?im_w=1200'),
(3, 'https://a0.muscache.com/im/pictures/c653fd86-090b-4456-85a1-82bfc2bfda76.jpg?im_w=1200'),
(3, 'https://a0.muscache.com/im/pictures/5506f206-8042-4b13-9239-59334d403a0d.jpg?im_w=1200'),
(3, 'https://a0.muscache.com/im/pictures/33462634-e146-4ada-8e2c-76fb1e84c420.jpg?im_w=1200'),
(3, 'https://a0.muscache.com/im/pictures/0c4d9ba7-2bc0-4086-98ed-ed5848fe8a21.jpg?im_w=1200'),
(4, 'https://a0.muscache.com/im/pictures/miso/Hosting-5904771/original/ab9a30d4-a6cf-4b3a-8416-cf7314ed5432.jpeg?im_w=1200'),
(4, 'https://a0.muscache.com/im/pictures/f2a6d022-23fc-4701-b703-806870cd9ed7.jpg?im_w=1440'),
(4, 'https://a0.muscache.com/im/pictures/miso/Hosting-5904771/original/411c07b5-c03b-4bf7-a0fc-65ea439be452.jpeg?im_w=1440'),
(4, 'https://a0.muscache.com/im/pictures/40036a49-6fc1-45f4-9988-9f49c26c902a.jpg?im_w=1440'),
(4, 'https://a0.muscache.com/im/pictures/miso/Hosting-5904771/original/e8429c29-d29b-4333-ae9f-4163265429fc.jpeg?im_w=1440'),
(5, 'https://a0.muscache.com/im/pictures/0d7b6751-98a9-4769-ab3d-63c6c3237cc3.jpg?im_w=1200'),
(5, 'https://a0.muscache.com/im/pictures/c187f9c5-fa8c-47ca-9fdd-0655c6566189.jpg?im_w=1200'),
(5, 'https://a0.muscache.com/im/pictures/f606a560-fb75-4a01-ab15-d256d8d300a9.jpg?im_w=1200'),
(5, 'https://a0.muscache.com/im/pictures/0edda9e2-4963-42b4-8f60-1a7d488d70f1.jpg?im_w=1200'),
(5, 'https://a0.muscache.com/im/pictures/d291acc4-bf06-4c59-bc27-d85a287fc0b4.jpg?im_w=1200'),
(6, 'https://a0.muscache.com/im/pictures/miso/Hosting-783090950382595560/original/55140ab8-9bd9-4bc1-848a-b6ff4fec1f49.jpeg?im_w=1200' ),
(6, 'https://a0.muscache.com/im/pictures/miso/Hosting-783090950382595560/original/f9766492-5f7b-4500-a9a9-5c7ed0def0f6.jpeg?im_w=1440' ),
(6,'https://a0.muscache.com/im/pictures/miso/Hosting-783090950382595560/original/422e9b80-ee75-4066-b310-b26c73783764.jpeg?im_w=1440' ),
(6, 'https://a0.muscache.com/im/pictures/miso/Hosting-783090950382595560/original/77f0165e-66d5-4451-8539-9325bd326fe4.jpeg?im_w=1440'),
(6,'https://a0.muscache.com/im/pictures/miso/Hosting-783090950382595560/original/058e7c3c-2f90-4e66-9d5d-1353a0681533.jpeg?im_w=1440' ),
(7,'https://a0.muscache.com/im/pictures/miso/Hosting-612559258605916614/original/b66a6286-ba97-4ffb-9483-d8d1a61354f0.jpeg?im_w=1200'),
(7,'https://a0.muscache.com/im/pictures/miso/Hosting-611907426138854889/original/8dcc8261-6806-4c6a-85b2-08dd1c171378.jpeg?im_w=1440'),
(7,'https://a0.muscache.com/im/pictures/miso/Hosting-611907426138854889/original/9f31b50e-6919-45a8-b4e0-e86cb605cb07.jpeg?im_w=1440'),
(7,'https://a0.muscache.com/im/pictures/miso/Hosting-611907426138854889/original/f0d9acef-ecc0-4401-b5d5-a31681cd6378.jpeg?im_w=1440'),
(7,'https://a0.muscache.com/im/pictures/miso/Hosting-611907426138854889/original/a5192ad8-ee61-40af-b28f-1b8f04c9eef2.jpeg?im_w=1440'),
(8,'https://a0.muscache.com/im/pictures/5c2fdb65-3da1-46d0-91dc-d7220e33f797.jpg?im_w=1200'),
(8,'https://a0.muscache.com/im/pictures/bd087b8a-eaa5-4cd4-98d8-86369753759a.jpg?im_w=1200'),
(8,'https://a0.muscache.com/im/pictures/5c2fdb65-3da1-46d0-91dc-d7220e33f797.jpg?im_w=1200'),
(8,'https://a0.muscache.com/im/pictures/9c131f99-2e10-4373-a9ef-874d6439c0db.jpg?im_w=1200'),
(8,'https://a0.muscache.com/im/pictures/07a90a6e-4e71-4191-82b4-5bffb75e9326.jpg?im_w=1200'),
(9 ,'https://a0.muscache.com/im/pictures/5db70666-78af-493c-8997-01c0575f19c3.jpg?im_w=1440'),
(9, 'https://a0.muscache.com/im/pictures/734bd0b3-3b4b-42b1-b44f-81e24864bda6.jpg?im_w=1200'),
( 9, 'https://a0.muscache.com/im/pictures/e7b3a179-d175-490c-9e9a-0d10f7afbee3.jpg?im_w=1440'),
( 9, 'https://a0.muscache.com/im/pictures/1fef7697-74d5-4e6c-b45f-658521c5405d.jpg?im_w=1440'),
(10, 'https://a0.muscache.com/im/pictures/46743116/d898b0b6_original.jpg?im_w=1200'),
(10, 'https://a0.muscache.com/im/pictures/46743193/3ff694fc_original.jpg?im_w=1440'),
(11 ,'https://a0.muscache.com/im/pictures/aafe8515-ee0e-4850-a738-cf2aed747405.jpg?im_w=720'),
(11, 'https://a0.muscache.com/im/pictures/e2ac5e70-ced6-4d6b-9b03-a0031172ccc4.jpg?im_w=1200'),
(11, 'https://a0.muscache.com/im/pictures/6b69708e-390e-478e-8080-ed68b8817a22.jpg?im_w=1440'),
(11, 'https://a0.muscache.com/im/pictures/miso/Hosting-53549739/original/2df3a52e-6ca0-4d74-bf40-059d40e33804.jpeg?im_w=1440'),
(12 ,'https://a0.muscache.com/im/pictures/666eeff6-12f5-4a94-a2d7-fc727c95e5cf.jpg?im_w=1440'),
(12, 'https://a0.muscache.com/im/pictures/miso/Hosting-4694846/original/f1404137-8d55-4b27-a18a-e689302178c2.jpeg?im_w=1200'),
(12,'https://a0.muscache.com/im/pictures/miso/Hosting-4694846/original/e38e4660-04fc-4e78-a792-7b9b6952cac8.jpeg?im_w=1440'),
(12, 'https://a0.muscache.com/im/pictures/miso/Hosting-4694846/original/687e24c5-6873-4976-8a21-fb653c2accff.jpeg?im_w=1440'),
(13, 'https://a0.muscache.com/im/pictures/3ca8c58e-2cac-431a-ae68-dc609127ecf3.jpg?im_w=1200'),
(13, 'https://a0.muscache.com/im/pictures/da4b24a0-f847-488f-8ebe-9a0aa8146b00.jpg?im_w=1440'),
(13, 'https://a0.muscache.com/im/pictures/e9cc3c26-d7af-489d-829f-1c930114cfb1.jpg?im_w=1200'),
(13, 'https://a0.muscache.com/im/pictures/a3627d6d-1366-4fc3-8976-e54369f8cc6d.jpg?im_w=1200'),
(13, 'https://a0.muscache.com/im/pictures/5f0fac15-3f51-4569-9ced-75e37b9e5409.jpg?im_w=1200'),
(14, 'https://a0.muscache.com/im/pictures/b1a3d1c1-d3cc-4dda-bec6-4d8e4e6d4822.jpg?im_w=1440'),
(14, 'https://a0.muscache.com/im/pictures/785b8961-97a4-4709-ba83-fad35ef4560c.jpg?im_w=1200'),
(14, 'https://a0.muscache.com/im/pictures/f0d7bfcc-ec41-41df-9b0b-b24729f9ce4c.jpg?im_w=1440'),
(14, 'https://a0.muscache.com/im/pictures/1c005d35-7e9b-4fd6-9991-ecd38365e1e0.jpg?im_w=1440'),
(14, 'https://a0.muscache.com/im/pictures/62498408-849a-4a04-93ca-3c35abf29e66.jpg?im_w=1440'),
(15, 'https://a0.muscache.com/im/pictures/miso/Hosting-46665306/original/a8277635-638c-4a8f-a472-a1ce6a6dd301.jpeg?im_w=1200'),
(15, 'https://a0.muscache.com/im/pictures/miso/Hosting-46665306/original/0770ee8c-2bb7-454b-815a-e8da87581c59.jpeg?im_w=1200'),
(15, 'https://a0.muscache.com/im/pictures/miso/Hosting-46665306/original/42a04016-5b11-4bd5-8c19-c6099ab58de0.jpeg?im_w=720'),
(15, 'https://a0.muscache.com/im/pictures/4b5c22fc-387c-4778-bfdc-142ab2ff0191.jpg?im_w=720'),
(15, 'https://a0.muscache.com/im/pictures/miso/Hosting-46665306/original/772ba654-23ad-494d-b736-fc14539da2bb.jpeg?im_w=1200'),
(16, 'https://a0.muscache.com/im/pictures/monet/Luxury-28540017/original/42cde801-3346-41c6-a63a-40bd5a2d5867?im_w=720'),
(16, 'https://a0.muscache.com/im/pictures/60b844b5-0e39-4b5b-88a6-389a6bb56a5e.jpg?im_w=1200'),
(16, 'https://a0.muscache.com/im/pictures/2a7fe38e-185d-46c3-903b-c8005b270ca5.jpg?im_w=1200'),
(16, 'https://a0.muscache.com/im/pictures/71d00c5c-82a6-4592-a49f-c4a8c5e22ef7.jpg?im_w=1200'),
(16, 'https://a0.muscache.com/im/pictures/24fc1fd5-44b4-4cdb-a53b-3f452201d6e1.jpg?im_w=720'),
(17, 'https://a0.muscache.com/im/pictures/prohost-api/Hosting-28765303/original/b86e5fa1-bd03-47fc-a98c-d86a4270782d.jpeg?im_w=1200'),
(17, 'https://a0.muscache.com/im/pictures/prohost-api/Hosting-28765303/original/56d48262-0e4a-4b57-bc19-ee80fb5eb321.jpeg?im_w=1200'),
(17, 'https://a0.muscache.com/im/pictures/prohost-api/Hosting-28765303/original/e0eac00d-b63f-466a-8942-57e7428cc4d7.jpeg?im_w=720'),
(17, 'https://a0.muscache.com/im/pictures/prohost-api/Hosting-28765303/original/6d7799c6-e328-4a55-b6e1-c86aff61cf1a.jpeg?im_w=1200'),
(17, 'https://a0.muscache.com/im/pictures/prohost-api/Hosting-28765303/original/86ad51d4-af47-4af0-9cb3-177914f6441e.jpeg?im_w=720'),
(18, 'https://a0.muscache.com/im/pictures/e3b5a0cf-03f2-483b-8c3a-c4ce16beb7cd.jpg?im_w=1200'),
(18, 'https://a0.muscache.com/im/pictures/dfe828a7-921c-48f0-abdb-aebf8ee6f7f7.jpg?im_w=1200'),
(18, 'https://a0.muscache.com/im/pictures/c9210953-2b53-4b10-9886-34271315c853.jpg?im_w=1200'),
(18, 'https://a0.muscache.com/im/pictures/a61ed86b-441d-47d7-9f1f-be3d5cbf84e5.jpg?im_w=720'),
(18, 'https://a0.muscache.com/im/pictures/0f06fcce-13cc-44ea-9b4e-966a9ab81ab3.jpg?im_w=1200'),
(19, 'https://a0.muscache.com/im/pictures/14152ff7-28fa-48cc-9c90-ac787fb5bb6b.jpg?im_w=1200'),
(19, 'https://a0.muscache.com/im/pictures/5f927ac6-1030-4eea-b34d-0da5bebcdaad.jpg?im_w=1200'),
(19, 'https://a0.muscache.com/im/pictures/a940b3cc-3f2e-4ba1-9fda-0c82cbd4b37a.jpg?im_w=720'),
(19, 'https://a0.muscache.com/im/pictures/4d6a08d6-cdc6-45f2-b582-439361274f08.jpg?im_w=1200'),
(19, 'https://a0.muscache.com/im/pictures/8619574d-a507-4df4-b55e-2b2c16d99893.jpg?im_w=1200'),
(20, 'https://a0.muscache.com/im/pictures/da1a1f70-5d49-4811-b5b9-d28d48f84fc7.jpg?im_w=1200'),
(20, 'https://a0.muscache.com/im/pictures/a54076b6-86a2-4d3f-8ee5-965712a8b240.jpg?im_w=1200'),
(20, 'https://a0.muscache.com/im/pictures/16760ffe-d134-47ab-96c4-bb8fc8d0684d.jpg?im_w=720'),
(20, 'https://a0.muscache.com/im/pictures/ac461f31-4046-4535-b7e5-01dda7c10b40.jpg?im_w=1200'),
(20, 'https://a0.muscache.com/im/pictures/dd200f0c-d7aa-4110-8b82-c05d01b150f4.jpg?im_w=720');



INSERT INTO house_details (max_guest, bedrooms, beds, baths, house_facilities, house_services, house_rules, bathrooms_facilities, house_id)
VALUES
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 1
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 2
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 3
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 4
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 5
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 6
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 7
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 8
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 9
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 10
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 11
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 12
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 13
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 14
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 15
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 16
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 17
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 18
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 19
),
(15, 5,5, 5, 
 'Elevator, 24-hour room service, Restaurant, Room service, Safety deposit box, WiFi in public area, TV',
 'Bellhop, Welcoming drinks, Concierge, Money changer, Doorman, 24-hour security',
 'Operational Hours	From 14:00 - Before 12:00',
 'Separate shower and tub, Shower, Bathrobe, Bathtub', 
 20
);




INSERT INTO reservations(house_id, user_id, check_in, check_out, total_price, status_id, booking_code, expired) VALUES
(1, 3,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823657ed727fdsad66', '2023-01-03 19:42:34.872514+07'),
(2, 4,'2023-02-03', '2023-02-04', 5600000, 2 , '7f810786-c257-4c6e-8ab3-823657eoihni1xxx6', '2023-01-03 19:42:34.872514+07'),
(3, 5,'2023-02-03', '2023-02-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823hdsdga9ni1daq6', '2023-01-03 19:42:34.872514+07'),
(4, 5,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823hoooooooovfso6', '2023-01-03 19:42:34.872514+07'),
(1, 5,'2023-02-03', '2023-02-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823hdsdga9xqqani16', '2023-01-03 19:42:34.872514+07'),
(6, 7,'2023-01-13', '2023-01-24', 9000000, 2 , '7f810786-c257-4c6e-8ab3-823hotrwooooo6iktyt', '2023-01-03 19:42:34.872514+07'),
(7, 8,'2023-01-13', '2023-01-24', 9000000, 2 , '7f810786-c257-4c6e-8ab3-823hotrwopppo6557', '2023-01-03 19:42:34.872514+07'),
(8, 9,'2023-01-13', '2023-01-24', 9000000, 2 , '7f810786-c257-4c6e-8ab3-823hotxxusooo6kky', '2023-01-03 19:42:34.872514+07'),
(9, 5,'2023-01-13', '2023-01-24', 9000000, 2 , '7f810786-c257-4c6e-8ab3-823hotxxuso788oo6', '2023-01-03 19:42:34.872514+07'),
(10, 3,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823657ed727f6632323', '2023-01-03 19:42:34.872514+07'),
(11, 8,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823657ed2775x122e26', '2023-01-03 19:42:34.872514+07'),
(12, 9,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823657ed2775x6rwqrq', '2023-01-03 19:42:34.872514+07'),
(13, 3,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823657ed2775x6ewqrq', '2023-01-03 19:42:34.872514+07'),
(14, 5,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823657ed2775x6rqqq', '2023-01-03 19:42:34.872514+07'),
(15, 4,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823657ed2775x6ddcxx', '2023-01-03 19:42:34.872514+07'),
(16, 4,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823657ed2732ffw6', '2023-01-03 19:42:34.872514+07'),
(17, 5,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-823657ed2775x6112', '2023-01-03 19:42:34.872514+07'),
(18, 7,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-8232fed2775x6ffs', '2023-01-03 19:42:34.872514+07'),
(19, 8,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-82365wd2775x6', '2023-01-03 19:42:34.872514+07'),
(20, 9,'2023-01-03', '2023-01-04', 3600000, 2 , '7f810786-c257-4c6e-8ab3-82365teed2775x6', '2023-01-03 19:42:34.872514+07');


INSERT INTO transactions(reservation_id,  house_id,user_id) VALUES
(1, 1, 3),
(2, 2, 4),
(3, 3, 5),
(4, 4, 5),
(5, 1, 5),
(6, 6, 7),
(7, 7, 8),
(8, 8, 9),
(9, 9, 5),
(10, 10, 3),
(11, 11, 8),
(12, 12, 9),
(13, 13, 3),
(14, 14, 5),
(15, 15, 4),
(16, 16, 4),
(17, 17, 5),
(18, 18, 7),
(19, 19, 8),
(20, 20, 9);
