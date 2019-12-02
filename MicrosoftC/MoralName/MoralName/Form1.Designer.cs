namespace MoralName
{
    partial class autoname
    {
        /// <summary>
        /// 必需的设计器变量。
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// 清理所有正在使用的资源。
        /// </summary>
        /// <param name="disposing">如果应释放托管资源，为 true；否则为 false。</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows 窗体设计器生成的代码

        /// <summary>
        /// 设计器支持所需的方法 - 不要修改
        /// 使用代码编辑器修改此方法的内容。
        /// </summary>
        private void InitializeComponent()
        {
            System.Windows.Forms.DataGridViewCellStyle dataGridViewCellStyle1 = new System.Windows.Forms.DataGridViewCellStyle();
            System.Windows.Forms.DataGridViewCellStyle dataGridViewCellStyle2 = new System.Windows.Forms.DataGridViewCellStyle();
            System.Windows.Forms.DataGridViewCellStyle dataGridViewCellStyle3 = new System.Windows.Forms.DataGridViewCellStyle();
            System.ComponentModel.ComponentResourceManager resources = new System.ComponentModel.ComponentResourceManager(typeof(autoname));
            this.birthday = new System.Windows.Forms.Label();
            this.birthDayInput = new System.Windows.Forms.DateTimePicker();
            this.firstNameInput = new System.Windows.Forms.TextBox();
            this.username = new System.Windows.Forms.Label();
            this.sex = new System.Windows.Forms.Label();
            this.namegrid = new System.Windows.Forms.DataGridView();
            this.namecol = new System.Windows.Forms.DataGridViewTextBoxColumn();
            this.button1 = new System.Windows.Forms.Button();
            this.label1 = new System.Windows.Forms.Label();
            this.label2 = new System.Windows.Forms.Label();
            this.limitName = new System.Windows.Forms.TextBox();
            this.limitType = new System.Windows.Forms.ComboBox();
            this.panel1 = new System.Windows.Forms.Panel();
            this.single_name = new System.Windows.Forms.RadioButton();
            this.more_name = new System.Windows.Forms.RadioButton();
            this.panel2 = new System.Windows.Forms.Panel();
            this.girl = new System.Windows.Forms.RadioButton();
            this.boy = new System.Windows.Forms.RadioButton();
            this.bottom_panel = new System.Windows.Forms.Panel();
            this.copyright_panel = new System.Windows.Forms.Panel();
            this.center_panel = new PanleModel.PanelEnhanced();
            this.title_sici = new System.Windows.Forms.Label();
            this.title_score = new System.Windows.Forms.Label();
            ((System.ComponentModel.ISupportInitialize)(this.namegrid)).BeginInit();
            this.panel1.SuspendLayout();
            this.panel2.SuspendLayout();
            this.center_panel.SuspendLayout();
            this.SuspendLayout();
            // 
            // birthday
            // 
            this.birthday.AutoSize = true;
            this.birthday.BackColor = System.Drawing.Color.Transparent;
            this.birthday.CausesValidation = false;
            this.birthday.Cursor = System.Windows.Forms.Cursors.No;
            this.birthday.Font = new System.Drawing.Font("仿宋", 15F, System.Drawing.FontStyle.Bold, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.birthday.ForeColor = System.Drawing.Color.SeaGreen;
            this.birthday.Location = new System.Drawing.Point(64, 153);
            this.birthday.Name = "birthday";
            this.birthday.Size = new System.Drawing.Size(93, 20);
            this.birthday.TabIndex = 0;
            this.birthday.Text = "出生日期";
            // 
            // birthDayInput
            // 
            this.birthDayInput.CalendarForeColor = System.Drawing.SystemColors.ButtonShadow;
            this.birthDayInput.Location = new System.Drawing.Point(173, 150);
            this.birthDayInput.MaxDate = new System.DateTime(2059, 12, 25, 23, 59, 0, 0);
            this.birthDayInput.MinDate = new System.DateTime(2019, 1, 1, 0, 0, 0, 0);
            this.birthDayInput.Name = "birthDayInput";
            this.birthDayInput.Size = new System.Drawing.Size(200, 21);
            this.birthDayInput.TabIndex = 1;
            // 
            // firstNameInput
            // 
            this.firstNameInput.ForeColor = System.Drawing.SystemColors.Highlight;
            this.firstNameInput.Location = new System.Drawing.Point(173, 27);
            this.firstNameInput.Name = "firstNameInput";
            this.firstNameInput.Size = new System.Drawing.Size(200, 21);
            this.firstNameInput.TabIndex = 2;
            // 
            // username
            // 
            this.username.AutoSize = true;
            this.username.BackColor = System.Drawing.Color.Transparent;
            this.username.Font = new System.Drawing.Font("仿宋", 15F, System.Drawing.FontStyle.Bold, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.username.ForeColor = System.Drawing.Color.SeaGreen;
            this.username.Location = new System.Drawing.Point(64, 25);
            this.username.Name = "username";
            this.username.Size = new System.Drawing.Size(95, 20);
            this.username.TabIndex = 3;
            this.username.Text = "姓    氏";
            // 
            // sex
            // 
            this.sex.AutoSize = true;
            this.sex.BackColor = System.Drawing.Color.Transparent;
            this.sex.Font = new System.Drawing.Font("仿宋", 15F, System.Drawing.FontStyle.Bold, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.sex.ForeColor = System.Drawing.Color.SeaGreen;
            this.sex.Location = new System.Drawing.Point(64, 200);
            this.sex.Name = "sex";
            this.sex.Size = new System.Drawing.Size(95, 20);
            this.sex.TabIndex = 4;
            this.sex.Text = "性    别";
            // 
            // namegrid
            // 
            this.namegrid.AllowUserToAddRows = false;
            this.namegrid.AllowUserToDeleteRows = false;
            this.namegrid.AllowUserToResizeColumns = false;
            this.namegrid.AllowUserToResizeRows = false;
            dataGridViewCellStyle1.BackColor = System.Drawing.Color.Honeydew;
            this.namegrid.AlternatingRowsDefaultCellStyle = dataGridViewCellStyle1;
            this.namegrid.BackgroundColor = System.Drawing.SystemColors.ControlLightLight;
            this.namegrid.CausesValidation = false;
            this.namegrid.CellBorderStyle = System.Windows.Forms.DataGridViewCellBorderStyle.SingleHorizontal;
            this.namegrid.ColumnHeadersBorderStyle = System.Windows.Forms.DataGridViewHeaderBorderStyle.Single;
            dataGridViewCellStyle2.Alignment = System.Windows.Forms.DataGridViewContentAlignment.MiddleLeft;
            dataGridViewCellStyle2.BackColor = System.Drawing.Color.DarkGreen;
            dataGridViewCellStyle2.Font = new System.Drawing.Font("仿宋", 8F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            dataGridViewCellStyle2.ForeColor = System.Drawing.SystemColors.Window;
            dataGridViewCellStyle2.SelectionBackColor = System.Drawing.SystemColors.Highlight;
            dataGridViewCellStyle2.SelectionForeColor = System.Drawing.SystemColors.HighlightText;
            dataGridViewCellStyle2.WrapMode = System.Windows.Forms.DataGridViewTriState.True;
            this.namegrid.ColumnHeadersDefaultCellStyle = dataGridViewCellStyle2;
            this.namegrid.ColumnHeadersHeight = 25;
            this.namegrid.ColumnHeadersHeightSizeMode = System.Windows.Forms.DataGridViewColumnHeadersHeightSizeMode.DisableResizing;
            this.namegrid.Columns.AddRange(new System.Windows.Forms.DataGridViewColumn[] {
            this.namecol});
            this.namegrid.Cursor = System.Windows.Forms.Cursors.Default;
            this.namegrid.EnableHeadersVisualStyles = false;
            this.namegrid.Font = new System.Drawing.Font("仿宋", 8F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.namegrid.GridColor = System.Drawing.SystemColors.ActiveCaption;
            this.namegrid.Location = new System.Drawing.Point(818, 1);
            this.namegrid.Name = "namegrid";
            this.namegrid.ReadOnly = true;
            this.namegrid.RowHeadersVisible = false;
            this.namegrid.RowTemplate.Height = 23;
            this.namegrid.ScrollBars = System.Windows.Forms.ScrollBars.Vertical;
            this.namegrid.SelectionMode = System.Windows.Forms.DataGridViewSelectionMode.FullRowSelect;
            this.namegrid.ShowCellErrors = false;
            this.namegrid.Size = new System.Drawing.Size(103, 481);
            this.namegrid.TabIndex = 17;
            this.namegrid.CellClick += new System.Windows.Forms.DataGridViewCellEventHandler(this.namegrid_CellClick);
            // 
            // namecol
            // 
            dataGridViewCellStyle3.BackColor = System.Drawing.SystemColors.Control;
            dataGridViewCellStyle3.Font = new System.Drawing.Font("宋体", 10.5F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.namecol.DefaultCellStyle = dataGridViewCellStyle3;
            this.namecol.HeaderText = "姓名";
            this.namecol.Name = "namecol";
            this.namecol.ReadOnly = true;
            // 
            // button1
            // 
            this.button1.Location = new System.Drawing.Point(173, 242);
            this.button1.Name = "button1";
            this.button1.Size = new System.Drawing.Size(75, 23);
            this.button1.TabIndex = 8;
            this.button1.Text = "提交";
            this.button1.UseVisualStyleBackColor = true;
            this.button1.Click += new System.EventHandler(this.button1_Click);
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.BackColor = System.Drawing.Color.Transparent;
            this.label1.Font = new System.Drawing.Font("仿宋", 15F, System.Drawing.FontStyle.Bold, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.label1.ForeColor = System.Drawing.Color.SeaGreen;
            this.label1.Location = new System.Drawing.Point(64, 69);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(95, 20);
            this.label1.TabIndex = 9;
            this.label1.Text = "字    数";
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.BackColor = System.Drawing.Color.Transparent;
            this.label2.Font = new System.Drawing.Font("仿宋", 10F, System.Drawing.FontStyle.Bold, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.label2.ForeColor = System.Drawing.Color.SeaGreen;
            this.label2.Location = new System.Drawing.Point(64, 111);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(98, 14);
            this.label2.TabIndex = 12;
            this.label2.Text = "特定字(选填)";
            // 
            // limitName
            // 
            this.limitName.Enabled = false;
            this.limitName.ForeColor = System.Drawing.SystemColors.HotTrack;
            this.limitName.Location = new System.Drawing.Point(173, 110);
            this.limitName.Name = "limitName";
            this.limitName.Size = new System.Drawing.Size(150, 21);
            this.limitName.TabIndex = 13;
            // 
            // limitType
            // 
            this.limitType.Enabled = false;
            this.limitType.FormattingEnabled = true;
            this.limitType.Items.AddRange(new object[] {
            "居中",
            "末尾"});
            this.limitType.Location = new System.Drawing.Point(326, 111);
            this.limitType.Name = "limitType";
            this.limitType.Size = new System.Drawing.Size(47, 20);
            this.limitType.TabIndex = 14;
            // 
            // panel1
            // 
            this.panel1.Controls.Add(this.single_name);
            this.panel1.Controls.Add(this.more_name);
            this.panel1.Location = new System.Drawing.Point(173, 69);
            this.panel1.Name = "panel1";
            this.panel1.Size = new System.Drawing.Size(200, 23);
            this.panel1.TabIndex = 15;
            // 
            // single_name
            // 
            this.single_name.AutoSize = true;
            this.single_name.Checked = true;
            this.single_name.Location = new System.Drawing.Point(7, 3);
            this.single_name.Name = "single_name";
            this.single_name.Size = new System.Drawing.Size(47, 16);
            this.single_name.TabIndex = 18;
            this.single_name.TabStop = true;
            this.single_name.Text = "单字";
            this.single_name.UseVisualStyleBackColor = true;
            this.single_name.CheckedChanged += new System.EventHandler(this.single_name_CheckedChanged);
            // 
            // more_name
            // 
            this.more_name.AutoSize = true;
            this.more_name.Location = new System.Drawing.Point(149, 3);
            this.more_name.Name = "more_name";
            this.more_name.Size = new System.Drawing.Size(47, 16);
            this.more_name.TabIndex = 19;
            this.more_name.Text = "多字";
            this.more_name.UseVisualStyleBackColor = true;
            this.more_name.CheckedChanged += new System.EventHandler(this.more_name_CheckedChanged_2);
            // 
            // panel2
            // 
            this.panel2.Controls.Add(this.girl);
            this.panel2.Controls.Add(this.boy);
            this.panel2.Location = new System.Drawing.Point(173, 197);
            this.panel2.Name = "panel2";
            this.panel2.Size = new System.Drawing.Size(200, 23);
            this.panel2.TabIndex = 16;
            // 
            // girl
            // 
            this.girl.AutoSize = true;
            this.girl.Location = new System.Drawing.Point(149, 4);
            this.girl.Name = "girl";
            this.girl.Size = new System.Drawing.Size(35, 16);
            this.girl.TabIndex = 8;
            this.girl.Text = "女";
            this.girl.UseVisualStyleBackColor = true;
            // 
            // boy
            // 
            this.boy.AutoSize = true;
            this.boy.Checked = true;
            this.boy.Location = new System.Drawing.Point(7, 3);
            this.boy.Name = "boy";
            this.boy.Size = new System.Drawing.Size(35, 16);
            this.boy.TabIndex = 7;
            this.boy.TabStop = true;
            this.boy.Text = "男";
            this.boy.UseVisualStyleBackColor = true;
            // 
            // bottom_panel
            // 
            this.bottom_panel.BackColor = System.Drawing.Color.Snow;
            this.bottom_panel.Location = new System.Drawing.Point(1, 332);
            this.bottom_panel.Name = "bottom_panel";
            this.bottom_panel.Size = new System.Drawing.Size(811, 150);
            this.bottom_panel.TabIndex = 18;
            this.bottom_panel.Visible = false;
            // 
            // copyright_panel
            // 
            this.copyright_panel.Location = new System.Drawing.Point(1, 485);
            this.copyright_panel.Name = "copyright_panel";
            this.copyright_panel.Size = new System.Drawing.Size(920, 24);
            this.copyright_panel.TabIndex = 19;
            this.copyright_panel.Paint += new System.Windows.Forms.PaintEventHandler(this.copyright_panel_Paint);
            // 
            // center_panel
            // 
            this.center_panel.BackColor = System.Drawing.Color.Transparent;
            this.center_panel.BorderStyle = System.Windows.Forms.BorderStyle.FixedSingle;
            this.center_panel.Controls.Add(this.title_sici);
            this.center_panel.Controls.Add(this.title_score);
            this.center_panel.ForeColor = System.Drawing.SystemColors.ControlLightLight;
            this.center_panel.Location = new System.Drawing.Point(398, 1);
            this.center_panel.Name = "center_panel";
            this.center_panel.Size = new System.Drawing.Size(414, 325);
            this.center_panel.TabIndex = 17;
            this.center_panel.Visible = false;
            // 
            // title_sici
            // 
            this.title_sici.AutoSize = true;
            this.title_sici.BackColor = System.Drawing.Color.SeaGreen;
            this.title_sici.Font = new System.Drawing.Font("宋体", 14.25F, ((System.Drawing.FontStyle)((System.Drawing.FontStyle.Bold | System.Drawing.FontStyle.Italic))), System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.title_sici.ForeColor = System.Drawing.Color.Linen;
            this.title_sici.Location = new System.Drawing.Point(-5, 136);
            this.title_sici.Name = "title_sici";
            this.title_sici.Size = new System.Drawing.Size(111, 19);
            this.title_sici.TabIndex = 1;
            this.title_sici.Text = "[名人名句]";
            // 
            // title_score
            // 
            this.title_score.AutoSize = true;
            this.title_score.BackColor = System.Drawing.Color.SeaGreen;
            this.title_score.Font = new System.Drawing.Font("宋体", 14.25F, ((System.Drawing.FontStyle)((System.Drawing.FontStyle.Bold | System.Drawing.FontStyle.Italic))), System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.title_score.ForeColor = System.Drawing.Color.Linen;
            this.title_score.Location = new System.Drawing.Point(-5, 7);
            this.title_score.Name = "title_score";
            this.title_score.Size = new System.Drawing.Size(111, 19);
            this.title_score.TabIndex = 0;
            this.title_score.Text = "[综合评分]";
            // 
            // autoname
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.BackgroundImage = global::MoralName.Properties.Resources._91;
            this.BackgroundImageLayout = System.Windows.Forms.ImageLayout.Stretch;
            this.ClientSize = new System.Drawing.Size(920, 507);
            this.Controls.Add(this.copyright_panel);
            this.Controls.Add(this.namegrid);
            this.Controls.Add(this.bottom_panel);
            this.Controls.Add(this.center_panel);
            this.Controls.Add(this.panel2);
            this.Controls.Add(this.panel1);
            this.Controls.Add(this.limitType);
            this.Controls.Add(this.limitName);
            this.Controls.Add(this.label2);
            this.Controls.Add(this.label1);
            this.Controls.Add(this.button1);
            this.Controls.Add(this.sex);
            this.Controls.Add(this.username);
            this.Controls.Add(this.firstNameInput);
            this.Controls.Add(this.birthDayInput);
            this.Controls.Add(this.birthday);
            this.Cursor = System.Windows.Forms.Cursors.Default;
            this.DoubleBuffered = true;
            this.FormBorderStyle = System.Windows.Forms.FormBorderStyle.FixedDialog;
            this.Icon = ((System.Drawing.Icon)(resources.GetObject("$this.Icon")));
            this.MaximizeBox = false;
            this.MinimizeBox = false;
            this.Name = "autoname";
            this.SizeGripStyle = System.Windows.Forms.SizeGripStyle.Hide;
            this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen;
            this.Text = "启名";
            ((System.ComponentModel.ISupportInitialize)(this.namegrid)).EndInit();
            this.panel1.ResumeLayout(false);
            this.panel1.PerformLayout();
            this.panel2.ResumeLayout(false);
            this.panel2.PerformLayout();
            this.center_panel.ResumeLayout(false);
            this.center_panel.PerformLayout();
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Label birthday;
        private System.Windows.Forms.DateTimePicker birthDayInput;
        private System.Windows.Forms.TextBox firstNameInput;
        private System.Windows.Forms.Label username;
        private System.Windows.Forms.Label sex;
        private System.Windows.Forms.Button button1;
        private System.Windows.Forms.Label label1;
        private System.Windows.Forms.Label label2;
        private System.Windows.Forms.TextBox limitName;
        private System.Windows.Forms.ComboBox limitType;
        private System.Windows.Forms.Panel panel1;
        private System.Windows.Forms.RadioButton single_name;
        private System.Windows.Forms.RadioButton more_name;
        private System.Windows.Forms.Panel panel2;
        private System.Windows.Forms.RadioButton girl;
        private System.Windows.Forms.RadioButton boy;
        private System.Windows.Forms.DataGridView namegrid;
        private PanleModel.PanelEnhanced center_panel;
        private System.Windows.Forms.Panel bottom_panel;
        private System.Windows.Forms.Label title_score;
        private System.Windows.Forms.Label title_sici;
        private System.Windows.Forms.DataGridViewTextBoxColumn namecol;
        private System.Windows.Forms.Panel copyright_panel;
    }
}

