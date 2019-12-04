using System;
using System.Collections.Generic;
using System.Text;
using Newtonsoft.Json;

namespace MoralName
{
    class JsonUtil
    {
        public static T DeserializeObject<T>(string json)
        {
            try
            {
                return JsonConvert.DeserializeObject<T>(json);
            }
            catch (Exception e){
                Console.WriteLine(e.Message);
            }
            return default(T);
           
        }

        public static string SerializeObject(object obj)
        {
            return JsonConvert.SerializeObject(obj);
        }
    }
}
