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
(4, 'WAITING CONFIRMATION')


INSERT INTO pickup_status VALUES 
(1, 'Pending Admin'),
(2, 'Awaiting Check-in Date'),
(3, 'On the Way Pickup'),
(4, 'On the Way Reservation'),
(5, 'Completed')

INSERT INTO users (email, fullname, password, address, role, city_id) VALUES 
('admin@gmail.com', 'admin giwang', '$2a$10$itR/9copEoG1cNdaMJfhDOMDGOf6yBmTxOBfWbc7PrqSHrZEn4/kC ', 'Perum Baros', 'admin', '1'),
('host@gmail.com', 'host giwang', '$2a$10$itR/9copEoG1cNdaMJfhDOMDGOf6yBmTxOBfWbc7PrqSHrZEn4/kC ', 'jl. Masjid Darussalam No. 11', 'host', '2'),
('user@gmail.com', 'user giwang', '$2a$10$itR/9copEoG1cNdaMJfhDOMDGOf6yBmTxOBfWbc7PrqSHrZEn4/kC ', 'jl. Masjid Yerussalem No. 11', 'user', '4'),
('giwangdk@gmail.com', 'user 1', '$2a$10$itR/9copEoG1cNdaMJfhDOMDGOf6yBmTxOBfWbc7PrqSHrZEn4/kC ', 'jl. Soeku Yerussalem No. 21', 'user', '3'),
('gidwikintan@gmail.com', 'user 2', '$2a$10$itR/9copEoG1cNdaMJfhDOMDGOf6yBmTxOBfWbc7PrqSHrZEn4/kC ', 'jl. Penjaittan Yerussalem No. 01', 'user', '6'),
('host2@gmail.com', 'host2 giwang', '$2a$10$itR/9copEoG1cNdaMJfhDOMDGOf6yBmTxOBfWbc7PrqSHrZEn4/kC ', 'jl. Masjid Darussalam No. 11', 'host', '10')



INSERT INTO wallets(balance, user_id) VALUES
( 1000000, 1),( 1000000, 2),( 1000000, 3),( 1000000, 4),( 1000000, 5), ( 100000, 6)

INSERT INTO games (user_id, chance, total_games_played) VALUES
( 1, 10, 0),
( 2, 10, 0),
( 3, 10, 0),
( 4, 10, 0),
( 5, 10, 0),
( 6, 10, 0)


INSERT INTO houses (name, city_id, user_id, price, description, location) VALUES 
(
 'Aura House 2bds Eco Bamboo House, Pool, River View', 
 11,
 2,
3600000, 
'Aura house is a beautiful & unique eco bamboo house built on the west bank of the River Ayung facing east to catch sunrise. Aura House is situated 25min away from Ubud, and 35min away from Canggu.',
'Canggu'),
(
 'Lespoir II - ocean view+huge pool+ butler', 
 11,
 2,
5600000, 
'Villa Lespoir II - Open on 20th Feb. 2020, our 2nd villa which is just about 40m away from Villa Lespoir I. Again luxury PANORAMIC view villa in tranquil northern Bali. It is at where the green vineyards and rice fields meet the ocean.',
'Seririt'),
(
 'Villa Saluka - Romantic & Serene with Sunset View', 
 12,
 6,
3600000, 
'Dear future guests, thank you so much for coming across our page. A little bit about Villa Saluka: the name is from my lovely little family. SAtria (my husband), LUise (me), KAi (my son). I initially built this house for us, but for a few reasons, we couldn’t live there. ',
 'Bantul'),
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
 'Banguntapan'),
(
 'Luxury 1 bedroom at parahyangan residence', 
4,
 6,
3600000, 
'Kick back and relax in this calm, stylish space.',
 'Cidadap'),
(
 'Brand New Bright Studio @ Landmark City Centre', 
 25,
 2,
2600000, 
'Our bright and lively brand new studio apartment at lv 12 in Landmark Residence is located strategically in the heart of Bandung. You will be close to everything when you stay at this centrally-located place. The complex is one of Bandungs most verdant and prestigious residence. ',
 'Canggu'),
(
 'Oribu Kidul, Direct Sunrise view of Mount Gede', 
 26,
 2,
4000000, 
'Relax, unwind, or even working. This tiny villa is a place for all. 
 Nestled in the feet of Mount Salak overviewing Mount Pangrango and our own garden. 
 Your solution for short getaway from citys hustle, 70 minutes away from Jakarta.',
 'Cigombong'),
 (
 'Villa Pusat kota Batu 2br,3menit dari Alun2 Batu', 
 9,
 6,
100000, 
'
Gaya industrial minimalis',
 'Bitu'),
 (
 'RiceField Villa Bukit Lawang', 
 2,
 6,
150000, 
'
It is a house of stone, bamboo and wood. Built 100% by hand with local materials in 2015. It is in the middle of a rice plantation in the rural area of Bukit Lawang and just 10 minutes walk from the bus station. Perfect for relaxing as a couple or with friends. It has mountain, sunset and sunrise views. A unique and authentic house in Bukit Lawang',
 'Bohorok'),
(
 'New Home with own beach frontage', 
 26,
 6,
150000, 
'
It has four Bedrooms and 2.5 Bathrooms and two large living areas opening onto huge decks to allow for relaxed outdoor beach-side living. Furnishing and decor is  modern and minimalist. Fully air-conditioned, fans, WiFi, TV and sound system.',
  'Pelabuhan Ratu'),
(
 'Quiet Villa and Pool with Mt. and Lake Batur View', 
 11,
 6,
6500000,
'This villa is located in Batur Kintamani Bali. Located right at the foot of Mount Batur so you can enjoy the beauty of Mount Batur and Lake Batur. The villa environment is very comfortable, the atmosphere is very calm, and the air is still fresh and far from the crowds of the city. You can enjoy the beautiful views of the mountains and lakes while swimming.',
  'Bangli'),
(
 '180° VIEW, PRIVATE POOL VILLA..', 
 5,
 2,
3500000,
'The villa sits in an elevated postion in one of the last remaining untouched parts of the island.',
  'Buyu')




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
(12, 'https://a0.muscache.com/im/pictures/miso/Hosting-4694846/original/687e24c5-6873-4976-8a21-fb653c2accff.jpeg?im_w=1440')