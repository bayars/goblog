---
date: "209-12-05"
draft: false
lastmod: "2019-12-05"
publishdate: "2019-12-05"
tags:
- linux
- CPU-algorithm
- algorithms
- replacement
title: Replacement Algoritmasi
---

# Replacement Algoritması:

Cache algoritmaları komutların veya algoritmaların optimize edilmesi, cache algoritmaları ile verilerin saklanması işleminin verimli olmasını sağlamak için kullanılır. Cache’in temiz tutulmasının önemi verilerin daha verimli bir şekilde saklanması ve bu sayede çalışma zamanı süresini kısaltması amaçlanmaktadır.  

Replacement algoritmaları yalnızca associative ve set associative ihtiyaçları vardır.  

### 1-) Least Recently Used (LRU):
Az sıklıkla kullanılan cache bellekteki verinin sıralanıp en az olanın çıkarılmasıdır. Peki bir veri bloğunun az kullanıldığının nasıl bilmektedir. Bu algoritma da her cache line’a bir age bits atanmaktadır ve age bits üzerinden hangi cache line’ın ne sıklıkla kullanılmakta olduğunu belirtir. Age bitleri sürekli takip edilmektedir. Age bitlerin durumuna göre az kullanılan veri cache bellekten silinir.  


### 2-) First-in-First-out (FIFO) :  
Veri yapılarında ki  queue’nun ilk giren ilk çıkar mantığının benzeridir. Cache bellek içinde ilk olan verinin yer ihtiyacı olduğunda ilk veriyle replacement edilmesidir.  

### 3-) Last-in-First-out (LIFO) :  
Veri yapılarında ki stack yapısının son giren ilk çıkar mantığının benzeridir. Cache bellekte en son gelmiş verinin replacement edilmesidir.  

### 4-) Time aware least recently used (TLRU) :   
Burada cachede saklanan her veriye geçerli yaşam süresi atanır. Bu veriler verilerin tutulduğu cache bellek üzerinde saklanır. Genellikle Information-centric networking (ICN) (Bili merkezli ağlar) , Content Delivery Networks (CDNs)(İçerik teslimatlı ağlar) ve genellikle gömülü ağ sistemlerinde bulunur. Time to use (TTU) sayesinde bir ağda sayfaya veya içeriğe ulaşma sıklığına göre cache bellekte saklanıp, sık kullanılmayan sayfayı ise memory de tutulup erişim süresini kısaltmaktadır.  

### 5-) Most recently used (MRU) :  
LRU’nun tersidir. Cache’de bulunan eski verilerin ne kadar eskiyse o kadar çok kullanılmıştır ve kullanılacaktır mantığıyla ortaya çıkmıştır. Bunu yapan kişiler LRU dan daha verimli olduğunu söylemektedir.  

### 6-) Random Replacement:  
Cache bellekten blokların rastgele seçilip replacement edilme işlemidir. Eğer  cache bellekte ki veriler geçiçi olarak veya bulunan verilerin önemi yoksa kullanılabilir.  

### 7-) Least Frequently Used (LFU) :  
Kullanılma sıklığına göre replacement yapılır. En küçük olan değerlerin çıkarılma işlemidir. Bu işlemler counter tutularak yapılmaktadır. Counterların değerleriyle de kullanılma sıklığı çıkarılmaktadır.  

### 8-) Adaptive Replacement Cache (ARC) :  
IBM tarafından geliştirilmiştir. LRU ve LFU’nun optimize edilmiş halidir.  

### 9-) Clock with adaptive replacement (CAR) :  
ARC’ye ek olarak clock tutulmaktadır. Kullanıcıların parametrelerine bağlı değildir. Kendi parametre kümesini kullanarak yapar. ARC’den daha verimlidir. Toplamda 4 tane double link list kullanır. Bunların ikisi clock verilerini, ikisi LRU listelerini belirler. T1, yenilik veya kısa zamanlı fayda. T2, Sıklık veya uzun zamanlı fayda. LRU listelerini bu verilerin değerlerine göre düzenlemektedir.  

### 10-) Pannier: Container-based caching algorithm for compound objects :   
Konteynır yapısının kullanıldığı sistemlerde kullanılır. Konteyner yapısında farklı, heterojen veya ayrı sistem verilerini ayırıp önceliklerine göre sıralar. Hangi konteyner daha fazla aktif kullanılıyorsa bu konteyner verileri daha fazla cache kullanabilir. Verileri hot data ve cold data diye ikiye ayırmıştır. Bunun için birden fazla aşama kullanmaktadır.  



* cache line size: 8 den 64 byte arasında veri boyutu  
** Direct : Doğrudan erişimli bellek  
*** Set associative: Doğrudan dönüşümlü cache bellek  
**** Associative: Çağrışımlı cache bellek  
