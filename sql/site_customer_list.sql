SELECT 
    dbo.jcusf01.cocode,
    dbo.jcusf01.custmast,
    dbo.jcusf01.custnum,
    dbo.jcusf01.startdate,
    dbo.jcusf01.taxpcnt,
    dbo.jcusf01.taxpcnt2,
    dbo.jcusf01.taxexempt,
    dbo.jcusf01.penalty,
    dbo.jcusf01.po_num,
    dbo.jcusf01.d_waiver,
    dbo.jcusf01.sitename,
    dbo.jcusf01.siteaddr,
    dbo.jcusf01.siteaddr2,
    dbo.jcusf01.sitecity,
    dbo.jcusf01.sitestate,
    dbo.jcusf01.sitezip,
    dbo.jcusf01.sitephone,    
    dbo.jcusf01.sitefax,
    dbo.jcusf01.sitezip4,
    dbo.jcusf01.sitecntry,
    dbo.jcusf01.super,
    
    dbo.jcusf07.bllmast,
    dbo.jcusf07.blladdr,
    dbo.jcusf07.blladdr2,
    dbo.jcusf07.bllcity,
    dbo.jcusf07.bllemail1,
    dbo.jcusf07.bllemail2,    
    dbo.jcusf07.bllfax,
    dbo.jcusf07.bllname,
    dbo.jcusf07.bllphone,
    dbo.jcusf07.bllstate,
    dbo.jcusf07.bllzip,
    dbo.jcusf07.bllcocode,
    dbo.jcusf07.bllcontact,    
    dbo.jcusf07.bllcountry,
                                
    dbo.jcusf09.acctstatus,
    dbo.jcusf09.accttype,
    dbo.jcusf09.county,
    dbo.jcusf09.custcode1,
    dbo.jcusf09.terms,    
    dbo.jcusf09.email,
    dbo.jcusf09.email2,
    dbo.jcusf09.emaillst,
    dbo.jcusf09.nomail,
    dbo.jcusf09.nomailing,

    dbo.jcusf01.FKjcustmast

FROM dbo.jcusf01 

INNER JOIN dbo.jcusf07 ON dbo.jcusf01.FKjcustmast = dbo.jcusf07.FKjcustmast 

INNER JOIN dbo.jcusf09 ON dbo.jcusf01.custnum = dbo.jcusf09.custnum