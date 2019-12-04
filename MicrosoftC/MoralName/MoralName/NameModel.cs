using System;
using System.Collections.Generic;
using System.Text;

namespace NameEntry
{
    public class UserInfo
    {

        public string firstName { get; set; }

        public string suffixName { get; set; }

        public string gender { get; set; }

        public string year { get; set; }

        public string month { get; set; }

        public string day { get; set; }


    }

    public class PreUser
    {
        public string firstName { get; set; }

        public string limitWord { get; set; }

        public string limitType { get; set; }

        public string gender { get; set; }

        public string single { get; set; }

    }

    public class UserResult
    {
        public string resultstatus { get; set; }
        public string errorinfo { get; set; }
        public string[] numbers { get; set; }

        public string[] verseArr { get; set; }

        public string[] detailArr { get; set; }
    }


    public class PreUserResult
    {
        public string resultstatus { get; set; }

        public string errorinfo { get; set; }

        public string[] list { get; set; }
    }

}

public class LinkNumbers
{
    //QQ
    public string qq { get; set; }

    //QQ群
    public string qqg { get; set; }

    //微信
    public string weixin { get; set; }

    //微信群
    public string weixing{ get; set;} 
}
