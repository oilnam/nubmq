Goal: Outperform rabbitmq in terms of pure throughput without sacrificing reliability in any way 

Issues:

- CONCURRENT SET OPERATIONS ARE FUCKING SLOW EVEN IN DIFFERENT SHARDS

- WE DONT RESIZE A LOT, BUT WHEN WE FUCKING DO, IT TAKES FOREVER, DONT FUCKING RESIZE WHEN YOU MUST, DO IT IN BACKGROUD WHEN YOU'RE FUCKING FREE AND THEN JUST REPLACE POINTER OR SMTH BRUH

- Resizes are fucking expensive, minimize the number of times we resize, it locks things down ffs

- SETs are left hanging when we resize, problem


