using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Net;
using System.IO;
using System.Runtime.Serialization.Json;
using NameEntry;

namespace MoralName
{
    public partial class autoname : Form
    {
        public autoname()
        {
            InitializeComponent();
        }

        private void addGridLine()
        {
            string firstName = this.firstNameInput.Text;
            string limitType = this.limitType.SelectedIndex.ToString();
            if (limitType == "0")
            {
                limitType = "center";
            }
            else
            {
                limitType = "end";
            }
            string limitWord = this.limitName.Text;
            string gender = "1";//男
            if (this.girl.Checked)
            {
                gender = "2";
            }
            bool haslimit = this.more_name.Checked;
            if (firstName == "")
            {
                MessageBox.Show("请输入姓氏");
                return;
            }
            string jsonData = preuser2str(firstName, limitWord, limitType,gender, haslimit);
            string result = postApi("http://127.0.0.1:8099/api/name/grasp", jsonData);
            if (result == "")
            {
                MessageBox.Show("请求超时...");
                return;
            }
            PreUserResult preUser = str2preUserResult(result);
            if (preUser.resultstatus != "1")
            {
                MessageBox.Show(preUser.errorinfo);
                return;
            }

            this.namegrid.Rows.Clear();

            for (int i = 0; i < preUser.list.Length; i++)
            {
                int index = this.namegrid.Rows.Add();
                this.namegrid.Rows[index].Cells[0].Value = preUser.list[i];
            }
        }


        private void button1_Click(object sender, EventArgs e)
        {

            addGridLine();
        }


        private void more_name_CheckedChanged_2(object sender, EventArgs e)
        {
            if (this.more_name.Checked)
            {
                this.limitName.Enabled = true;
                this.limitType.Enabled = true;
                if (this.limitType.SelectedIndex == -1)
                {
                    this.limitType.SelectedIndex = 0;
                }
            }
        }

        private void single_name_CheckedChanged(object sender, EventArgs e)
        {
            if (this.single_name.Checked)
            {
                this.limitName.Enabled = false;
                this.limitType.Enabled = false;
            }
        }




        private void namegrid_CellClick(object sender, DataGridViewCellEventArgs e)
        {
            int index = e.RowIndex;
            string selectName = this.namegrid.Rows[index].Cells[0].Value.ToString();
            string firstName = this.firstNameInput.Text;
            string suffixName = selectName.Replace(firstName, "");

            string year = this.birthDayInput.Value.Year.ToString();
            string month = this.birthDayInput.Value.Month.ToString();
            string day = this.birthDayInput.Value.Day.ToString();
            string genderCode = "1";//默认为男
            if (this.girl.Checked)
            {
                genderCode = "2";
            }

            //TODO 调用接口
            string jsonData = user2str(firstName, suffixName, genderCode, year, month, day);

            string result = postApi("http://127.0.0.1:8099/api/name/parse", jsonData);
            if (result == "")
            {
                MessageBox.Show("请求超时...");
                return;
            }
            if (!this.center_panel.Visible)
            {
                this.center_panel.Visible = true;
            }

            UserResult userResult = str2userResult(result);

            if (userResult.resultstatus != "1")
            {
                MessageBox.Show(userResult.errorinfo);
            }

            score_insert(userResult.numbers);
            sici_insert(userResult.verseArr);
            bttom_panel_insert(userResult.detailArr);
        }


        private void score_insert(string[] scores)
        {
            this.center_panel.Controls.Clear();
            this.center_panel.Controls.Add(this.title_sici);
            this.center_panel.Controls.Add(this.title_score);

            string[] types = new string[] { "文化印象", "五行八字", "生　　肖", "五格数理" };
            for (int i = 0; i < 4; i++)
            {
                Label left_label1 = new Label();
                Label score_label1 = new Label();
                left_label1.AutoSize = false;
                left_label1.Width = 150;
                left_label1.Height = 20;
                left_label1.Text = types[i];
                left_label1.Parent = this.center_panel;
                left_label1.Location = new System.Drawing.Point(5, 35 + 25 * i);
                left_label1.Font = new System.Drawing.Font("新宋体", 10F, System.Drawing.FontStyle.Bold);
                left_label1.ForeColor = System.Drawing.Color.Green;

                score_label1.AutoSize = false;
                score_label1.Width = 30;
                score_label1.Height = 20;
                score_label1.Text = scores[i];
                score_label1.Font = new System.Drawing.Font("新宋体", 10F, System.Drawing.FontStyle.Bold);
                score_label1.Parent = this.center_panel;
                score_label1.Location = new System.Drawing.Point(180, 35 + 25 * i);
                score_label1.ForeColor = System.Drawing.Color.Green;

            }

        }

        private void sici_insert(string[] siciArr)
        {
            for (int i = 0; i < siciArr.Length; i++)
            {
                Label sici_label1 = new Label();
                sici_label1.Parent = this.center_panel;
                sici_label1.Width = 320;
                sici_label1.Height = 40;
                sici_label1.Location = new System.Drawing.Point(5, 160 + 40 * i);
                sici_label1.Font = new System.Drawing.Font("微软雅黑", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point);
                sici_label1.ForeColor = System.Drawing.Color.Sienna;
                if (siciArr[i].Contains("。") && siciArr[i].Contains("－"))
                {
                    sici_label1.Text = siciArr[i].Replace("。", "。\n");
                }
                else
                {
                    sici_label1.Text = siciArr[i];
                }

            }

        }

        private void bttom_panel_insert(string[] strs)
        {
            if (!this.bottom_panel.Visible)
            {
                this.bottom_panel.Visible = true;
            }
            else {
                this.bottom_panel.Controls.Clear();
            }

            for (int i = 0; i < strs.Length; i++)
            {
                Label content_label1 = new Label();
                content_label1.AutoSize = false;
                content_label1.Width = 700;
                content_label1.Height = 18;
                content_label1.Text = strs[i];
                content_label1.Parent = this.bottom_panel;
                content_label1.Location = new System.Drawing.Point(20, 3 + 18 * i);
            }
        }

        private string postApi(string url, string postData)
        {
            try
            {
                byte[] bytes = System.Text.Encoding.UTF8.GetBytes(postData);
                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(url);
                request.Method = "post";
                request.ContentType = "application/json;charset=UTF-8";
                request.ContentLength = bytes.Length;
                Stream requestStream = request.GetRequestStream();
                requestStream.Write(bytes, 0, bytes.Length);
                requestStream.Close();
                HttpWebResponse response = (HttpWebResponse)request.GetResponse();

                string json = null;
                using (StreamReader reader = new StreamReader(response.GetResponseStream(), System.Text.Encoding.GetEncoding("UTF-8")))
                {
                    json = reader.ReadToEnd();
                }
                return json;
            }
            catch
            {
                return "";
            }

        }

        private string user2str(string firstName, string suffixName, string genderCode, string year, string month, string day)
        {
            UserInfo user = new UserInfo();
            user.firstName = firstName;
            user.suffixName = suffixName;
            user.gender = genderCode;
            user.year = year;
            user.month = month;
            user.day = day;
            DataContractJsonSerializer serializer = new DataContractJsonSerializer(typeof(UserInfo));
            System.IO.MemoryStream ms = new MemoryStream();
            serializer.WriteObject(ms, user);
            System.IO.StreamReader reader = new StreamReader(ms);
            ms.Position = 0;
            string jsonData = reader.ReadToEnd();
            reader.Close();
            ms.Close();
            return jsonData;
        }

        private string preuser2str(string firstName, string limitWord, string limitType,string gender, bool haslimit)
        {
            PreUser preUser = new PreUser();
            if (haslimit)
            {
                preUser.limitWord = limitWord;
                preUser.limitType = limitType;
            }
            else
            {
                //单字查询
                preUser.single = "single";
            }
            preUser.firstName = firstName;
            preUser.gender = gender;
            DataContractJsonSerializer serializer = new DataContractJsonSerializer(typeof(PreUser));
            System.IO.MemoryStream ms = new MemoryStream();
            serializer.WriteObject(ms, preUser);
            System.IO.StreamReader reader = new StreamReader(ms);
            ms.Position = 0;
            string jsonData = reader.ReadToEnd();
            reader.Close();
            ms.Close();
            return jsonData;
        }


        private UserResult str2userResult(string str)
        {
            // 实例化DataContractJsonSerializer对象，需要待序列化的对象类型  
            DataContractJsonSerializer serializer = new DataContractJsonSerializer(typeof(UserResult));
            //把Json传入内存流中保存  
            MemoryStream stream = new MemoryStream(Encoding.UTF8.GetBytes(str));
            // 使用ReadObject方法反序列化成对象  
            object ob = serializer.ReadObject(stream);
            UserResult result = (UserResult)ob;
            return result;
        }

        private PreUserResult str2preUserResult(string str)
        {
            // 实例化DataContractJsonSerializer对象，需要待序列化的对象类型  
            DataContractJsonSerializer serializer = new DataContractJsonSerializer(typeof(PreUserResult));
            //把Json传入内存流中保存  
            MemoryStream stream = new MemoryStream(Encoding.UTF8.GetBytes(str));
            // 使用ReadObject方法反序列化成对象  
            object ob = serializer.ReadObject(stream);
            PreUserResult result = (PreUserResult)ob;
            return result;
        }

        private LinkNumbers str2linknumbers(string str)
        {
            // 实例化DataContractJsonSerializer对象，需要待序列化的对象类型  
            DataContractJsonSerializer serializer = new DataContractJsonSerializer(typeof(LinkNumbers));
            //把Json传入内存流中保存  
            MemoryStream stream = new MemoryStream(Encoding.UTF8.GetBytes(str));
            // 使用ReadObject方法反序列化成对象  
            object ob = serializer.ReadObject(stream);
            LinkNumbers result = (LinkNumbers)ob;
            return result;
        }


        private void copyright_panel_Paint(object sender, PaintEventArgs e)
        {

            string result = postApi("http://127.0.0.1:8099/api/name/links", "");
            if (result == "")
            {
                MessageBox.Show("连接服务器失败...");
                return;
            }

            LinkNumbers links = str2linknumbers(result);
            string text = "工具仅供个人交流学习使用，如需导出更多名字或了解更多信息欢迎咨询：";
            
            Label contentL = new Label();
            contentL.Text = text;
            contentL.Location = new System.Drawing.Point(100, 6);
            contentL.Name = "copyright_year";
            contentL.AutoSize = true;
            contentL.TabIndex = 3;
            contentL.Parent = this.copyright_panel;
            int widthpoint = 500;
            if (links.qqg != "")
            {
                Label qqglink = new Label();
                qqglink.Text = "QQ群: " + links.qqg;
                qqglink.Location = new System.Drawing.Point(widthpoint, 6);
                //qqglink.Font = new System.Drawing.Font("微软雅黑", 8F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point);
                qqglink.AutoSize = true;
                qqglink.Parent = this.copyright_panel;
                widthpoint += 8 * links.qqg.Length + 20;

            }
            else if (links.qq != "")
            {
                Label qqlink = new Label();
                qqlink.Text = "QQ: " + links.qq;
                qqlink.Location = new System.Drawing.Point(widthpoint, 6);
                //qqlink.Font = new System.Drawing.Font("微软雅黑", 8F);
                qqlink.AutoSize = true;
                qqlink.Parent = this.copyright_panel;
                widthpoint += 8 * links.qq.Length + 20;
            }

            if (links.weixing != "")
            {
                Label weixinglink = new Label();
                weixinglink.Text = "微信群: " + links.weixing;
                weixinglink.Location = new System.Drawing.Point(widthpoint, 6);
                weixinglink.Font = new System.Drawing.Font("微软雅黑", 8F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point);
                weixinglink.AutoSize = true;
                weixinglink.Parent = this.copyright_panel;
            }
            else if (links.weixin != "")
            {
                Label weixinlink = new Label();
                weixinlink.Text = "微信: " + links.weixin;
                weixinlink.Location = new System.Drawing.Point(widthpoint, 6);
                //weixinlink.Font = new System.Drawing.Font("微软雅黑", 8F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point);
                weixinlink.AutoSize = true;
                weixinlink.Parent = this.copyright_panel;
            }

        }
        
    }



}
