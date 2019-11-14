using System.Windows.Forms;

namespace PanleModel
{
    class PanelEnhanced : Panel
    {
        public PanelEnhanced()
        {
            SetStyle(ControlStyles.UserPaint | ControlStyles.AllPaintingInWmPaint | ControlStyles.OptimizedDoubleBuffer | ControlStyles.ResizeRedraw | ControlStyles.SupportsTransparentBackColor, true);
        }
    }
}
