SELECT 
    c1.cocode,
    c1.custmast,
    c1.custnum,
    c1.startdate,
    c1.taxpcnt,
    c1.taxpcnt2,
    c1.taxexempt,
    c1.penalty,
    c1.po_num,
    c1.d_waiver,
    c1.sitename,
    c1.siteaddr,
    c1.siteaddr2,
    c1.sitecity,
    c1.sitestate,
    c1.sitezip,
    c1.sitephone,    
    c1.sitefax,
    c1.sitezip4,
    c1.sitecntry,
    c1.super,
    
    c7.bllmast,
    c7.blladdr,
    c7.blladdr2,
    c7.bllcity,
    c7.bllemail1,
    c7.bllemail2,    
    c7.bllfax,
    c7.bllname,
    c7.bllphone,
    c7.bllstate,
    c7.bllzip,
    c7.bllcocode,
    c7.bllcontact,    
    c7.bllcountry,
                                
    c9.acctstatus,
    c9.accttype,
    c9.county,
    c9.custcode1,
    c9.terms,    
    c9.email,
    c9.email2,
    c9.emaillst,
    c9.nomail,
    c9.nomailing,

    c1.FKjcustmast

FROM jcusf01 as c1

INNER JOIN jcusf07 as c7 ON c1.FKjcustmast = c7.FKjcustmast 

INNER JOIN jcusf09 as c9 ON c1.custnum = c9.custnum